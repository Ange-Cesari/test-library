package testlibrary

// Version is the current version of the library
const Version = "0.1.0"

// NewClient creates a new client instance
func NewClient() *Client {
	return &Client{}
}

// Client provides functionality for test-library
type Client struct {
	// Add your fields here
}

// Hello returns a friendly greeting
func (c *Client) Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}

// ExampleFunction is an example function that can be called directly
// without creating a client instance
func ExampleFunction(input string) string {
	return "Example output for: " + input
}
