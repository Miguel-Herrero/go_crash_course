package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 5 {
		fmt.Println(i)
		// i = i +1
		i++
	}

	// Short method
	for f := 1; f <= 5; f++ {
		fmt.Printf("Number %d\n", f)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
