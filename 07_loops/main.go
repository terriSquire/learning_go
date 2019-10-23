package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz
	/* Loop from 1 - 100.
	For every multiple of 3 print "Fizz"
	For every multiple of 5 print "Buzz"
	For every multiple of 3 AND 5 print "FizzBuzz"
	*/
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
