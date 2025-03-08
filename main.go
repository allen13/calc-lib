package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/allen13/calc-lib/calc"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Calc!")
	fmt.Println("Enter mathematical expressions to evaluate (or 'quit' to exit)")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		input = strings.TrimSpace(input)

		if input == "quit" || input == "exit" {
			break
		}

		if input == "" {
			continue
		}

		result, err := calc.Evaluate(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("= %g\n", result)
		}
	}

	fmt.Println("Goodbye!")
}
