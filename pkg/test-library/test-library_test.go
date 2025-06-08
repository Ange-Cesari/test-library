package testlibrary

import (
	"testing"
)

func TestHello(t *testing.T) {
	client := NewClient()
	
	t.Run("with name", func(t *testing.T) {
		got := client.Hello("Test")
		want := "Hello, Test!"
		if got != want {
			t.Errorf("Hello() = %q, want %q", got, want)
		}
	})
	
	t.Run("without name", func(t *testing.T) {
		got := client.Hello("")
		want := "Hello, World!"
		if got != want {
			t.Errorf("Hello() = %q, want %q", got, want)
		}
	})
}

func TestExampleFunction(t *testing.T) {
	got := ExampleFunction("test")
	want := "Example output for: test"
	if got != want {
		t.Errorf("ExampleFunction() = %q, want %q", got, want)
	}
}

func TestCalculateSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"single value", []int{5}, 5},
		{"multiple values", []int{1, 2, 3, 4, 5}, 15},
		{"negative values", []int{-1, -2, 3, 4}, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CalculateSum(test.input)
			if result != test.expected {
				t.Errorf("Expected sum %d but got %d", test.expected, result)
			}
		})
	}
}
