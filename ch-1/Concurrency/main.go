package main

import (
	"fmt"
	"time"
)

// Simple function to print a message
func f() {
	fmt.Println("f function")
}

// Function that calculates the length of a string and sends it to a channel
func strlen(s string, c chan int) {
	c <- len(s) // Send the length of the string to channel 'c'
}

func main() {
	// ---------------------- GOROUTINE EXAMPLE ----------------------
	go f()                       // Call function 'f' in a new goroutine (runs concurrently)
	time.Sleep(1 * time.Second)  // Pause main for 1 second to allow goroutine to finish
	fmt.Println("main function") // Will likely print after "f function"

	// ---------------------- CHANNELS EXAMPLE ----------------------
	c := make(chan int) // Create a channel of type int

	// Launch two goroutines that send string lengths into the same channel
	go strlen("Salutations", c) // Will send 11
	go strlen("World", c)       // Will send 5

	// Receive values from the channel (order is not guaranteed)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) // Print the two values and their sum
}
