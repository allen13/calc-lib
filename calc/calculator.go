package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Calculator errors
var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionByZero    = errors.New("division by zero")
	ErrInvalidOperator   = errors.New("invalid operator")
	ErrInvalidToken      = errors.New("invalid token in expression")
)

// Evaluate parses and evaluates a mathematical expression string
func Evaluate(expression string) (float64, error) {
	// Remove all whitespace
	expression = strings.ReplaceAll(expression, " ", "")
	if expression == "" {
		return 0, ErrInvalidExpression
	}

	// Tokenize the expression
	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	// Process the tokens using a basic implementation of the Shunting Yard algorithm
	result, err := evaluateTokens(tokens)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// Token types
type tokenType int

const (
	numberToken tokenType = iota
	operatorToken
	leftParenToken
	rightParenToken
)

// Token represents a token in the expression
type token struct {
	tokenType tokenType
	value     string
}

// Tokenize the expression into numbers and operators
func tokenize(expression string) ([]token, error) {
	var tokens []token
	var currentNumber strings.Builder
	var hasDecimal bool

	for i, char := range expression {
		switch {
		case unicode.IsDigit(char):
			currentNumber.WriteRune(char)
		case char == '.':
			if hasDecimal {
				return nil, fmt.Errorf("%w: multiple decimal points", ErrInvalidToken)
			}
			hasDecimal = true
			currentNumber.WriteRune(char)
		case isOperator(char):
			// Flush any accumulated number
			if currentNumber.Len() > 0 {
				tokens = append(tokens, token{tokenType: numberToken, value: currentNumber.String()})
				currentNumber.Reset()
				hasDecimal = false
			}

			// Handle unary minus for negative numbers
			if char == '-' && (i == 0 || isOperator(rune(expression[i-1])) || expression[i-1] == '(') {
				currentNumber.WriteRune(char)
			} else {
				tokens = append(tokens, token{tokenType: operatorToken, value: string(char)})
			}
		case char == '(':
			// Flush any accumulated number
			if currentNumber.Len() > 0 {
				tokens = append(tokens, token{tokenType: numberToken, value: currentNumber.String()})
				currentNumber.Reset()
				hasDecimal = false
			}
			tokens = append(tokens, token{tokenType: leftParenToken, value: "("})
		case char == ')':
			// Flush any accumulated number
			if currentNumber.Len() > 0 {
				tokens = append(tokens, token{tokenType: numberToken, value: currentNumber.String()})
				currentNumber.Reset()
				hasDecimal = false
			}
			tokens = append(tokens, token{tokenType: rightParenToken, value: ")"})
		default:
			return nil, fmt.Errorf("%w: '%c'", ErrInvalidToken, char)
		}
	}

	// Flush any remaining number
	if currentNumber.Len() > 0 {
		tokens = append(tokens, token{tokenType: numberToken, value: currentNumber.String()})
	}

	return tokens, nil
}

// Check if a character is a valid operator
func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

// Get operator precedence
func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

// Evaluate an operation
func applyOperator(op string, b, a float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("%w: '%s'", ErrInvalidOperator, op)
	}
}

// Evaluate tokens using the Shunting Yard algorithm
func evaluateTokens(tokens []token) (float64, error) {
	var numbers []float64
	var operators []string

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		switch t.tokenType {
		case numberToken:
			num, err := strconv.ParseFloat(t.value, 64)
			if err != nil {
				return 0, fmt.Errorf("%w: '%s'", ErrInvalidToken, t.value)
			}
			numbers = append(numbers, num)
		case operatorToken:
			for len(operators) > 0 && operators[len(operators)-1] != "(" &&
				precedence(operators[len(operators)-1]) >= precedence(t.value) {
				
				if len(numbers) < 2 {
					return 0, ErrInvalidExpression
				}
				
				op := operators[len(operators)-1]
				operators = operators[:len(operators)-1]
				
				b := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				
				a := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				
				result, err := applyOperator(op, b, a)
				if err != nil {
					return 0, err
				}
				
				numbers = append(numbers, result)
			}
			operators = append(operators, t.value)
		case leftParenToken:
			operators = append(operators, t.value)
		case rightParenToken:
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				if len(numbers) < 2 {
					return 0, ErrInvalidExpression
				}
				
				op := operators[len(operators)-1]
				operators = operators[:len(operators)-1]
				
				b := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				
				a := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				
				result, err := applyOperator(op, b, a)
				if err != nil {
					return 0, err
				}
				
				numbers = append(numbers, result)
			}
			
			if len(operators) == 0 || operators[len(operators)-1] != "(" {
				return 0, ErrInvalidExpression
			}
			
			// Pop the left parenthesis
			operators = operators[:len(operators)-1]
		}
	}

	// Process any remaining operators
	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return 0, ErrInvalidExpression
		}
		
		if len(numbers) < 2 {
			return 0, ErrInvalidExpression
		}
		
		op := operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		
		b := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		
		a := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		
		result, err := applyOperator(op, b, a)
		if err != nil {
			return 0, err
		}
		
		numbers = append(numbers, result)
	}

	if len(numbers) != 1 {
		return 0, ErrInvalidExpression
	}

	return numbers[0], nil
}