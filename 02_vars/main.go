package main

import (
	"fmt"
	"strconv"
)

// Using var - can also declare/assign from within function (global variable)
var name = "Ally"

func main() {
	// Shorthand declaration & assignment - more concise but can ONLY
	// be used from within a function (e.g. scope issue)
	age := strconv.Itoa(44)

	fmt.Println("Hello there, " + name + "! Are you truly " + age + "?")
	fmt.Printf("%T\n", age)

	name2, email := "Cayman", "cayman_the_croc@sharpteeth.com"
	fmt.Println("Hey " + name2 + "! Is your email address still " + email + "?")

}
