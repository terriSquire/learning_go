package main

import (
	"fmt"
	"time"
)

func greeting(name string) string {
	return "Hello, " + name + "!"
}

func concatStr(string1, string2, string3 string) string {
	return string1 + string2 + string3
}

func main() {
	p := fmt.Println
	now := time.Now()
	p(greeting("Sal the Sloth!"))
	p(concatStr("How ", "Ya ", "Been?"))
	p(now)

}
