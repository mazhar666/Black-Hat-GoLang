package main

import "fmt"

// foo takes an empty interface and determines the underlying type using a type switch
func foo(i interface{}) {
	// Type switch to check the type of the interface value 'i'
	switch i.(type) {
	case int:
		fmt.Println("I'm an integer!")
	case string:
		fmt.Println("I'm a string!")
	default:
		fmt.Println("Unknown type!")
	}
}

func main() {
	// ---------------------- IF STATEMENT EXAMPLE ----------------------
	x := 10 // Declare an integer variable
	if x == 1 {
		fmt.Println("x is 1")
	} else {
		fmt.Println("X is not equal to 1") // This line will execute
	}

	// ---------------------- SWITCH STATEMENT EXAMPLE ----------------------
	y := "foo" // Declare a string variable
	switch y {
	case "foo":
		fmt.Println("Found foo") // This line will execute
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}

	// ---------------------- TYPE SWITCH EXAMPLE USING INTERFACE ----------------------
	foo(1)     // Output: I'm an integer!
	foo("bar") // Output: I'm a string!
	foo(true)  // Output: Unknown type!

	// ---------------------- STANDARD FOR LOOP ----------------------
	for i := 0; i < 10; i++ {
		fmt.Println(i) // Print numbers from 0 to 9
	}

	// ---------------------- RANGE LOOP OVER SLICE ----------------------
	nums := []int{2, 4, 6, 8} // A slice of integers
	for idx, val := range nums {
		fmt.Println(idx, val) // Print index and value
	}
}
