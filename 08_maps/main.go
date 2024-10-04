package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key-value pairs
	// emails["Colin"] = "colinting1001@gmail.com"
	// emails["John"] = "john@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Decalre map and add key-value pairs
	emails := map[string]string{"Colin": "colinting@gmail.com", "John": "john@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Colin"])

	// Delete from map
	delete(emails, "John")
	fmt.Println(emails)
}
