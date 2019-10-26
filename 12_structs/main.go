package main

import "fmt"

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Structs are important in Go
// Sort of like classes in other languages
// Pointer receiver and Value receiver methods
func main() {
	// Init person using struct
	person1 := Person{firstName: "Ellen", lastName: "Salamander",
		city: "Anytown", gender: "f", age: 1}

	// Alternative assignment (shortcut)
	person2 := Person{"Gregor", "Metamorpho", "Secret City", "m", 111}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person2.firstName)

}
