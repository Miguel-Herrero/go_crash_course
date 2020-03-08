package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// // Assign key:value
	// emails["Bob"] = "bob@golang.com"
	// emails["Mary"] = "mary@golang.com"

	// Define and Assign
	emails := map[string]string{"Bob": "bob@golang.com", "Mary": "mary@golang.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])
	delete(emails, "Bob")
	fmt.Println(emails)
}
