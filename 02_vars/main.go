// Variabes and types

package main

import "fmt"

func main() {
	// Shorthand method of decalaring variables
	name := "Tom"
	size := 1.3

	// Init multiple variables at once
	email, address := "tom@tom.com", "120 Seneca St, NY"

	// var name string = "Tom"
	var age int32 = 22
	const isActive = true

	fmt.Println(name, age, isActive, email, address)

	// Print typeof
	fmt.Printf("%T\n", size)
}
