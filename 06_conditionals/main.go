package main

import "fmt"

func main() {
	x := 5
	y := 10

	// If else
	if x < y || x > 0 {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Switch
	color := "green"
	switch color {
	case "red":
		fmt.Println("Color is red")
	default:
		fmt.Println("Color is unknown")
	}

}
