package main

import "errors"

// Define a custom error type based on string
type MyError string

// Implement the Error() method to satisfy the built-in error interface
func (e MyError) Error() string {
	return string(e) // Convert MyError (which is a string) back to a string
}

// Function that returns a standard error using errors.New
func foo() error {
	return errors.New("Some Error Occurred") // Creates a new error with a message
}

func main() {
	// Handle the error returned by foo
	if err := foo(); err != nil {
		println(err.Error()) // Print the error message if an error occurred
	}
}

// This is a re-declaration of Go’s built-in `error` interface for reference.
// Normally, you don’t need to declare this—it exists in the standard library.
type error interface {
	Error() string
}
