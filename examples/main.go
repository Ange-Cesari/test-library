package main

import (
	"encoding/json"
	"fmt"

	"github.com/Ange-Cesari/test-library"
)

func main() {
	// Demonstrate Calculate function
	sum := testlibrary.Calculate(10, 5)
	fmt.Printf("Calculate(10, 5) = %d\n", sum)

	// Demonstrate Multiply function
	product := testlibrary.Multiply(10, 5)
	fmt.Printf("Multiply(10, 5) = %d\n", product)

	// Demonstrate Greet function
	greeting := testlibrary.Greet("Developer")
	fmt.Println(greeting)

	// Demonstrate Info function
	info := testlibrary.Info()
	infoJSON, _ := json.MarshalIndent(info, "", "  ")
	fmt.Printf("System Information:\n%s\n", string(infoJSON))
}
