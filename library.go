// Package testlibrary provides utility functions for the test library.
package testlibrary

import (
	"fmt"
)

// Version of the library
const Version = "v0.1.0"

// Calculate adds two integers and returns the result
func Calculate(a, b int) int {
	return a + b
}

// Multiply multiplies two integers and returns the result
func Multiply(a, b int) int {
	return a * b
}

// Greet returns a greeting message
func Greet(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// Info returns information about the library
func Info() string {
	return fmt.Sprintf("test-library %s - A simple Go library for testing", Version)
}
