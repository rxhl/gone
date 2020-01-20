package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Tom"))
	fmt.Println(add(3, 4))
}
