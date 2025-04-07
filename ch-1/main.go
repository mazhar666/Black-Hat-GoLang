package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

func main() {
	var count = int(42)
	ptr := &count
	fmt.Println(*ptr)
	*ptr = 100
	fmt.Println(count)

	var guy = new(Person)
	guy.Name = "Dave"
	guy.SayHello()

}

type Friend interface {
	SayHello()
}
