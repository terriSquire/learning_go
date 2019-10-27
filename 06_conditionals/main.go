// package main

// import "fmt"

// func main() {
// 	x := 50
// 	y := 25

// 	// If else
// 	if x < y {
// 		fmt.Printf("%d is less than %d\n", x, y)
// 	} else {
// 		fmt.Printf("%d is less than %d\n", y, x)
// 	}

// 	// else if
// 	color := "teal"

// 	if color == "lavender" {
// 		fmt.Println("The color is lavender")
// 	} else if color == "indigo" {
// 		fmt.Println("The color is indigo")
// 	} else {
// 		fmt.Println("The color is neither lavender nor indigo")
// 	}

// 	// Switch
// 	switch color {
// 	case "lavender":
// 		fmt.Println("We've got lavender")
// 	case "indigo":
// 		fmt.Println("Here's indigo")
// 	default:
// 		fmt.Println("I have no idea what color this is")
// 	}
// }

/*
Another lesson in conditionals
starting from the beginning
*/

// This program reports whether a grade is passing or failing
package main

import (
	"bufio"
	"fmt"
	"log" // this package logs a msg to terminal & stops program in case of error
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter racer name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(name)

	fmt.Print("Enter racer rank: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	rank, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Need to declare "medal" outside the "if" statements,
	// so that it's still in scope afterwards.
	var medal string
	if rank == 1 {
		medal = "gold"
	} else if rank == 2 {
		medal = "silver"
	} else if rank == 3 {
		medal = "bronze"
	} else {
		medal = "participant"
	}

	fmt.Println(name, "gets a", medal, "medal!")
}
