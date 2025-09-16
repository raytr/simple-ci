package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/raytr/simple-ci/calculator"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: simple-ci <operation> <num1> <num2>")
		fmt.Println("Operations: add, subtract, multiply, divide")
		os.Exit(1)
	}

	operation := os.Args[1]
	num1, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Printf("Error parsing first number: %v\n", err)
		os.Exit(1)
	}

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Printf("Error parsing second number: %v\n", err)
		os.Exit(1)
	}

	var result float64
	var calcErr error

	switch operation {
	case "add":
		result = calculator.Add(num1, num2)
	case "subtract":
		result = calculator.Subtract(num1, num2)
	case "multiply":
		result = calculator.Multiply(num1, num2)
	case "divide":
		result, calcErr = calculator.Divide(num1, num2)
		if calcErr != nil {
			fmt.Printf("Error: %v\n", calcErr)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown operation: %s\n", operation)
		fmt.Println("Supported operations: add, subtract, multiply, divide")
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f\n", result)
}
