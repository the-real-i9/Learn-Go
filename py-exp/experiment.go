package main

import (
	"fmt"
	"pythontypes/pystring"
)

func main() {
	fmt.Printf("%q", pystring.Center("Hey", 10, ' '))
}
