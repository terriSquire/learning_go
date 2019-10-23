package main

import "fmt"

func main() {
	x := 50
	y := 25

	// If else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "teal"

	if color == "lavender" {
		fmt.Println("The color is lavender")
	} else if color == "indigo" {
		fmt.Println("The color is indigo")
	} else {
		fmt.Println("The color is neither lavender nor indigo")
	}

	// Switch
	switch color {
	case "lavender":
		fmt.Println("We've got lavender")
	case "indigo":
		fmt.Println("Here's indigo")
	default:
		fmt.Println("I have no idea what color this is")
	}
}
