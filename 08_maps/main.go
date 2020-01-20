package main

import "fmt"

func main() {
	// Hashmap
	emails := make(map[string]string)

	// Assign k-v
	emails["Bob"] = "bob@bob.com"
	emails["Alice"] = "alice@alice.com"
	emails["Sharon"] = "sharon@sharon.com"

	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Alice")
	fmt.Println(emails)

	// Shorthand
	users := map[string]string{"bob": "bob@gmail.com"}
	fmt.Println(users)
}
