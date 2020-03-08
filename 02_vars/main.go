package main

import "fmt"

func main() {
	// Using var
	var name = "John"
	var age int32 = 27
	const isCool = true

	// Shorthand
	size := 1.4
	email, country := "john@doe.com", "Brazil"

	fmt.Println(name, age, isCool, size, email, country)
	fmt.Printf("%T\n", size)
}
