package main

import "fmt"

func main() {
	// a pointer points to the memory address or location of a value that's
	// in a variable
	a := 5
	b := &a // assigns b to a pointer of a

	fmt.Println(a, b) // b prints the memory address of a
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) // prints *int bc it's a pointer

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a) // same thing as *b

	// Change val with pointer
	*b = 10
	fmt.Println(a) // note that the val of a has actually been changed to 10

	// pointers are useful when you need to pass a lot of data stored,
	// passing the address instead of actual data can increase performance
	// and speed
}
