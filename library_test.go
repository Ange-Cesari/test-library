package testlibrary

import (
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -2, -4, -6},
		{"mixed numbers", -5, 10, 5},
		{"zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Calculate(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Calculate(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 15},
		{"negative numbers", -2, -4, 8},
		{"mixed numbers", -5, 10, -50},
		{"zeros", 0, 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"regular name", "John", "Hello, John! Welcome to test-library."},
		{"empty name", "", "Hello, ! Welcome to test-library."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Greet(tt.input)
			if result != tt.expected {
				t.Errorf("Greet(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestInfo(t *testing.T) {
	info := Info()

	// Check that all expected keys are present
	requiredKeys := []string{"library_version", "go_version", "os", "architecture", "current_time"}
	for _, key := range requiredKeys {
		if _, ok := info[key]; !ok {
			t.Errorf("Info() missing required key: %q", key)
		}
	}

	// Check that library_version matches the expected constant
	if info["library_version"] != Version {
		t.Errorf("Info()[\"library_version\"] = %q; expected %q", info["library_version"], Version)
	}

	// Check that go_version starts with "go"
	if !strings.HasPrefix(info["go_version"], "go") {
		t.Errorf("Info()[\"go_version\"] = %q; expected to start with \"go\"", info["go_version"])
	}
}
