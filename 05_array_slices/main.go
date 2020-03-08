package main

import "fmt"

func main() {
	// Arrays
	// var fruits [2]string

	// fruits[0] = "Apple"
	// fruits[1] = "Orange"

	// fruits := [2]string{"Apple", "Orange"}

	// fmt.Println(fruits)

	// Slices
	fruitsSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitsSlice))
	fmt.Println(fruitsSlice[1:3])
}
