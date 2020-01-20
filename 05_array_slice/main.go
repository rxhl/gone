package main

import "fmt"

func main() {
	// Declare
	var fruitArr [2]string

	// Assign
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr[0]) // Apple

	// Declare + Assign
	fruits := [2]string{"Apple", "Orange"}
	fmt.Println(fruits)

	// Slice
	fruitSlice := []string{"Apple", "Orange", "Grapes"}
	fmt.Println(len(fruitSlice))

	// Range [startIndex, endIndex)
	fmt.Println(fruitSlice[1:2])
}
