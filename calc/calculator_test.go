package calc

import (
	"errors"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		expected float64
		wantErr  error
	}{
		{"Basic addition", "2+2", 4, nil},
		{"Basic subtraction", "5-3", 2, nil},
		{"Basic multiplication", "4*3", 12, nil},
		{"Basic division", "8/2", 4, nil},
		{"Complex expression", "3+4*2/(1-5)", 1, nil},
		{"Negative numbers", "-3+4", 1, nil},
		{"Decimal numbers", "3.5+2.1", 5.6, nil},
		{"Parentheses", "(2+3)*(4-1)", 15, nil},
		{"Multiple operations", "2+3*4-5", 9, nil},
		{"Multiple parentheses", "((2+3)*4)-5", 15, nil},
		{"Whitespace handling", " 3 + 4 * 2 ", 11, nil},
		{"Empty expression", "", 0, ErrInvalidExpression},
		{"Division by zero", "5/0", 0, ErrDivisionByZero},
		{"Invalid character", "5+a", 0, ErrInvalidToken},
		{"Invalid expression", "5++3", 0, ErrInvalidExpression},
		{"Unbalanced parentheses", "(2+3", 0, ErrInvalidExpression},
		{"Unbalanced closing", "2+3)", 0, ErrInvalidExpression},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Evaluate(tt.expr)
			
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Expected error %v, got nil", tt.wantErr)
					return
				}
				
				if !errors.Is(err, tt.wantErr) && err.Error() != tt.wantErr.Error() {
					t.Errorf("Expected error %v, got %v", tt.wantErr, err)
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}