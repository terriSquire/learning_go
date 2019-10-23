package main

import "fmt"

func main() {
	/*
		// Declare map (key value pair)
		emails := make(map[string]string)

		//Assign key values
		emails["Cheryl"] = "cheryl@she_shed.com"
		emails["Judy"] = "judy@she-she-er_she_shed.com"
	*/

	// Declare map and add key values (shortcut)
	emails := map[string]string{"Cheryl": "cheryl@she_shed.com",
		"Judy": "judy@she-she-er_she_shed.com"}

	// Add kv after map has been declared/assigned kv
	emails["Sharon"] = "sharon@roadtrip.com"

	fmt.Println(emails)
	fmt.Println(emails["Cheryl"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Cheryl")
	fmt.Println(emails)

}
