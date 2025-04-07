package main

import "fmt"

// Define a struct representing a person with Name and Age fields
type Person struct {
	Name string
	Age  int
}

// Define an empty struct representing a dog
type Dog struct{}

// Person implements the Friend interface by having a SayHello method
func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

// Define an interface named Friend that requires a SayHello method
type Friend interface {
	SayHello()
}

// Greet accepts anything that implements the Friend interface and calls its SayHello method
func Greet(f Friend) {
	f.SayHello()
}

// Dog also implements the Friend interface with its own SayHello method
func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

func main() {
	// Declare an integer variable and assign its address to a pointer
	var count = int(42)
	ptr := &count // pointer to count

	// Modify the value at the memory address pointed to by ptr
	*ptr = 100 // equivalent to count = 100

	// Create a new instance of Person using 'new'
	var guy = new(Person)
	guy.Name = "Dave"

	// Call the SayHello method directly on the person
	guy.SayHello()

	// Use the Greet function which takes any Friend type
	Greet(guy) // guy implements Friend

	// Create a new instance of Dog and pass it to Greet
	var dog = new(Dog)
	Greet(dog) // dog also implements Friend
}
