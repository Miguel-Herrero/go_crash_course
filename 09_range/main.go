package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	ids := []int{32, 43, 27, 11, 98}

	// Loop through ids
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
	fmt.Printf("Sum: %d\n", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@golang.com", "Mary": "mary@golang.com"}

	for key, value := range emails {
		fmt.Printf("Name: %s - Email: %s\n", key, value)
	}
}
