package main

import (
	"fmt"

	"github.com/Ange-Cesari/test-library"
)

func main() {
	// Display library information
	fmt.Println(testlibrary.Info())
	
	// Use the Calculate function
	result := testlibrary.Calculate(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	
	// Use the Multiply function
	product := testlibrary.Multiply(4, 7)
	fmt.Printf("4 * 7 = %d\n", product)
	
	// Use the Greet function
	greeting := testlibrary.Greet("Library User")
	fmt.Println(greeting)
}
