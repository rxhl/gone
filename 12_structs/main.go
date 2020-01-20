package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Methods for struct stay outside the struct body
// 1. Greeting method (Value receiver)
func (p Person) greet() string {
	return p.firstName + " is " + strconv.Itoa(p.age) + " years old."
}

// 2. Pointer receiver (when mutating a value)
func (p *Person) hasBday() {
	p.age++
}

func main() {
	// Init person using struct
	bob := Person{firstName: "Bob", lastName: "Dylan"}
	fmt.Println(bob)

	// Alternative init
	sam := Person{"Sam", "Smith", "Tokyo", "M", 20}
	fmt.Println(sam)

	sam.hasBday()
	fmt.Println(sam.greet())
}
