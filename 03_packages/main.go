package main

// Packages can either be imported via 'go get' command, or
// you can create your own and import using a similar path below

import (
	"fmt"

	"github.com/terriSquire/learning_go/03_packages/strutil"
)

func main() {
	fmt.Println(strutil.Reverse("!rotagilla ,retal uoy eeS"))
}
