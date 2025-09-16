package calculator

import (
	"errors"
	"math"
)

// Add performs addition of two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract performs subtraction of two numbers
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply performs multiplication of two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide performs division of two numbers
// Returns an error if attempting to divide by zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Power calculates a raised to the power of b
func Power(a, b float64) float64 {
	return math.Pow(a, b)
}

// Sqrt calculates the square root of a number
// Returns an error for negative inputs
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number is not allowed")
	}
	return math.Sqrt(a), nil
}

// IsEven checks if a number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// Abs returns the absolute value of a number
func Abs(a float64) float64 {
	return math.Abs(a)
}
