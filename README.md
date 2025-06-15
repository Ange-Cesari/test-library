# test-library

A simple Go library for testing, demonstrating library packaging with GoReleaser and Scotter.

## Features

- Go 1.21 compliant
- Simple calculation, text formatting, and info functions
- Semantic versioning with `v` prefix
- Comprehensive test coverage
- CI/CD pipelines with GitHub Actions
- Release automation with GoReleaser

## Installation

```bash
go get github.com/Ange-Cesari/test-library@v0.1.0
```

## Usage

```go
package main

import (
    "fmt"
    
    "github.com/Ange-Cesari/test-library"
)

func main() {
    // Print library information
    fmt.Println(testlibrary.Info())
    
    // Use calculation functions
    sum := testlibrary.Calculate(5, 3)
    fmt.Printf("5 + 3 = %d\n", sum)
    
    // Use greeting function
    greeting := testlibrary.Greet("User")
    fmt.Println(greeting)
}
```

## Available Functions

- `Calculate(a, b int) int` - Adds two integers
- `Multiply(a, b int) int` - Multiplies two integers
- `Greet(name string) string` - Returns a greeting message
- `Info() string` - Returns information about the library

## Development

This project follows the conventional commit specification. Please read the [COMMITS.md](./COMMITS.md) file for more information.
