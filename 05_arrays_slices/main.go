package main

import "fmt"

func main() {
	// Arrays need to be a fixed length (Slices are basically arrays without fixed lengths)
	var fosterDogs [2]string

	// Assign values
	fosterDogs[0] = "Benji"
	fosterDogs[1] = "Lassie"

	// Declare & Assign (shortcut)
	availableFosters := [2]string{"Banjo", "Lucky"}

	fmt.Println(fosterDogs)
	fmt.Println(availableFosters)

	// Slices
	fosterSlice := []string{"Aggie", "Bandit", "Carly", "Dougie", "Eugenie"}
	fmt.Println(fosterSlice)
	fmt.Println(len(fosterSlice))
	fmt.Println(fosterSlice[1:3])

}
