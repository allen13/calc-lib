# calc-lib

A simple calculator library for Go that supports basic arithmetic operations and parentheses.

## Installation

```bash
go get github.com/allen13/calc-lib
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/allen13/calc-lib/calc"
)

func main() {
    result, err := calc.Evaluate("2 + 3 * (4 - 1)")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Result: %g\n", result)
}
```

## Features

- Basic arithmetic operations (+, -, *, /)
- Parentheses support
- Proper operator precedence
- Error handling for invalid expressions
- Support for decimal numbers
- Handles negative numbers

## Error Handling

The library provides several error types:
- `ErrInvalidExpression`: Invalid mathematical expression
- `ErrDivisionByZero`: Division by zero attempted
- `ErrInvalidOperator`: Invalid operator encountered
- `ErrInvalidToken`: Invalid token in expression

## How It Works: Step-by-Step

This calculator uses the [Shunting Yard algorithm](https://en.wikipedia.org/wiki/Shunting_yard_algorithm) developed by Edsger Dijkstra to parse and evaluate mathematical expressions. Here's a breakdown of how it works:

### 1. Expression Tokenization (`tokenize` function)

When you call `calc.Evaluate("2 + 3 * 4")`:

1. First, all whitespace is removed for consistent parsing
2. The expression is broken down into tokens (numbers, operators, parentheses)
3. Each character is analyzed and converted to appropriate tokens:
   - Digits and decimal points are collected to form number tokens
   - Operators (+, -, *, /) become operator tokens
   - Parentheses are treated as special tokens
   - Special handling for negative numbers (unary minus)
4. The result is an array of tokens: `[2, +, 3, *, 4]`

### 2. Expression Evaluation (`evaluateTokens` function)

Once tokenized, the Shunting Yard algorithm is applied:

1. Two stacks are maintained:
   - A numbers stack for operands
   - An operators stack for operators and parentheses
   
2. For each token:
   - If it's a number: push it onto the numbers stack
   - If it's an operator:
     - While there are operators on the stack with higher or equal precedence, pop and apply them
     - Then push the current operator
   - If it's a left parenthesis: push it onto the operators stack
   - If it's a right parenthesis: evaluate everything back to the matching left parenthesis

3. Operator precedence is respected:
   - Multiplication (*) and division (/) have higher precedence than addition (+) and subtraction (-)
   - Operations within parentheses are evaluated first

4. When operators are applied:
   - Pop two numbers from the stack (operands)
   - Apply the operator
   - Push the result back onto the numbers stack

### 3. Example Evaluation of "2 + 3 * 4"

1. Push 2 onto the numbers stack: `numbers = [2]`
2. See +, push onto operators stack: `operators = [+]`
3. Push 3 onto the numbers stack: `numbers = [2, 3]`
4. See *, push onto operators stack: `operators = [+, *]`
5. Push 4 onto the numbers stack: `numbers = [2, 3, 4]`
6. End of expression, process remaining operators:
   - Pop * and apply it to 3 and 4, push result: `numbers = [2, 12]`
   - Pop + and apply it to 2 and 12, push result: `numbers = [14]`
7. Return the final number on the stack: 14

### 4. Handling Parentheses

For expressions with parentheses like "2 * (3 + 4)":

1. When a left parenthesis is encountered, it's pushed onto the operators stack
2. When a right parenthesis is encountered:
   - Operators are applied until the matching left parenthesis is found
   - The parentheses are then discarded
3. This ensures that operations within parentheses are evaluated first

### 5. Error Handling

The calculator carefully validates input and operations:
- Checks for invalid characters in expressions
- Prevents division by zero
- Validates proper parentheses matching
- Ensures the expression is well-formed

## License

MIT License 