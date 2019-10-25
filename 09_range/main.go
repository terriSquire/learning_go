package main

import "fmt"

func main() {
	// Loop through arrays, maps, slices
	ids := []int{19, 45, 44, 87, 49}

	// Loop thru ids with an index (i)
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index - instead of "for i, use for _,"
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum:", sum)

	// Range with map
	emails := map[string]string{"Cheryl": "cheryl@she_shed.com", "Judy": "judy@shesheersheshed.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v) // %s is a placeholder for k and v
	}

	// Get just the key
	for k := range emails {
		fmt.Printf("%s\n", k)
		// or
		fmt.Println("Name: " + k)
	}
}
