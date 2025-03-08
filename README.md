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

## License

MIT License 