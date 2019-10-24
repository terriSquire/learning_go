package main

import "fmt"

func main() {
	// Loop through arrays, maps, slices
	ids := []int{19, 45, 44, 87, 49}

	// Loop thru ids with an index (i)
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
}
