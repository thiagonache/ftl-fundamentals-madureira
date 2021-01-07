// Package calculator provides a library for simple calculations in Go.
package calculator

import "errors"

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes tow numbers and return the result
func Divide(a, b float64) (float64, error) {
	// return 2, errors.New("Anta não divide por 0 porra....")
	if b == 0 {
		return 0, errors.New("Anta não divide por 0 porra....")
	}
	return a / b, nil
}
