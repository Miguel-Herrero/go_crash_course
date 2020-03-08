package main

import "fmt"

func main() {
	a := 5
	var b *int = &a

	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)

	// Use * to read value from address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
