package main

import "fmt"

func main() {
	// Pass by reference is often cheaper

	a := 5
	b := &a

	// Address of a in hex
	fmt.Println(b)
	fmt.Printf("%T\n", b) // *int, integer pointer

	// Get the value from a known address
	fmt.Println(*b)

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}
