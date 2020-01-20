package main

import "fmt"

func main() {
	x := 5
	y := 10

	// If else
	if x <= y {
		fmt.Printf("%d is lte %d\n", x, y)
	} else {
		fmt.Printf("%d is gt %d\n", y, x)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Println("Red")
	} else if color == "blue" {
		fmt.Println("Blue")
	} else {
		fmt.Println("Not")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Not")
	}
}
