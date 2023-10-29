package main

import (
	"fmt"
	"learn/helpers" // an imported custom package (module)
)

func main() {
	var greeting = helpers.Greet("Bro")
	fmt.Print(greeting)
}
