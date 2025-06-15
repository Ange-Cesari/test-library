package testlibrary

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 5, 3, 8},
		{"Negative numbers", -2, -3, -5},
		{"Mixed numbers", -5, 3, -2},
		{"Zero values", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.a, tt.b); got != tt.expected {
				t.Errorf("Calculate(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 4, 5, 20},
		{"Negative numbers", -3, -4, 12},
		{"Mixed numbers", -2, 6, -12},
		{"Zero value", 7, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.a, tt.b); got != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
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
		{"With name", "John", "Hello, John!"},
		{"Empty name", "", "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.input); got != tt.expected {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestInfo(t *testing.T) {
	info := Info()
	if info == "" {
		t.Error("Info() returned empty string")
	}
}
