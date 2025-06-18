// Package testlibrary provides utility functions for demonstration purposes.
package testlibrary

import (
	"fmt"
	"runtime"
	"time"
)

// Version represents the current version of the library.
const Version = "0.1.0"

// Calculate adds two integers and returns the result.
func Calculate(a, b int) int {
	return a + b
}

// Multiply multiplies two integers and returns the result.
func Multiply(a, b int) int {
	return a * b
}

// Greet returns a greeting message with the provided name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to test-library.", name)
}

// Info returns information about the current system.
func Info() map[string]string {
	return map[string]string{
		"library_version": Version,
		"go_version":      runtime.Version(),
		"os":              runtime.GOOS,
		"architecture":    runtime.GOARCH,
		"current_time":    time.Now().Format(time.RFC3339),
	}
}
