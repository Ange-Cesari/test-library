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
