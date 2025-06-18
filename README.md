# test-library

A lightweight Go utility library providing common functions for demonstration purposes.

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
    // Basic calculations
    sum := testlibrary.Calculate(5, 3)
    fmt.Printf("5 + 3 = %d\n", sum)
    
    product := testlibrary.Multiply(5, 3)
    fmt.Printf("5 × 3 = %d\n", product)
    
    // Generate a greeting
    greeting := testlibrary.Greet("Developer")
    fmt.Println(greeting)
    
    // Get system information
    info := testlibrary.Info()
    for key, value := range info {
        fmt.Printf("%s: %s\n", key, value)
    }
}
```

## Features

* Basic arithmetic functions
* Greeting generator
* System information utilities
* Full test coverage
* Semantic versioning
* Continuous integration with GitHub Actions
* Release automation with GoReleaser

## Development

This project follows the conventional commit specification. Please read the [COMMITS.md](./COMMITS.md) file for more information.
