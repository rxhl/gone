package main

import "fmt"

func main() {
	// Long method
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Short method
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Test
	for i := 1; i <= 100; i++ {
		res := ""
		if i%3 == 0 {
			res += "Fizz"
		}
		if i%5 == 0 {
			res += "Buzz"
		}
		if res == "" {
			fmt.Println(i)
		} else {
			fmt.Println(res)
		}
	}
}
