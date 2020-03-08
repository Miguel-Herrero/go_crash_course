package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + ". My age is " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person := Person{firstName: "Jane", lastName: "Doe", city: "Madrid", gender: "f", age: 27}

	// Alternative
	person2 := Person{"John", "Smith", "Madrid", "m", 27}

	fmt.Println(person.firstName, person.lastName)

	fmt.Println(person.greet())

	person.hasBirthday()

	maridito := &person2
	person.getMarried(maridito.lastName)

	fmt.Println(person)

}
