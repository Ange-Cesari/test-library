package main

import (
	"fmt"

	"github.com/Ange-Cesari/test-library/pkg/test-library"
)

func main() {
	// Create a new client
	client := testlibrary.NewClient()
	
	// Use the client to say hello
	greeting := client.Hello("User")
	fmt.Println(greeting)
	
	// Use a direct function from the library
	result := testlibrary.ExampleFunction("test input")
	fmt.Println(result)
}
