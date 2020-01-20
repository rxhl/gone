package main

import "fmt"

func main() {
	ids := []int{20, 30}

	// Loop through IDs
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range with maps
	users := map[string]string{"bob": "bob@gmail.com", "Sharon": "Sharon@gmail.com"}
	for k, v := range users {
		fmt.Printf("%s: %s\n", k, v)
	}
}
