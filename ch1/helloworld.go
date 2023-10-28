package main

// this is a special type of package.
// A package different from "main" is a custom package
// The "main" package is where execution starts, and it must be the only "main" package among all other imported packages

import (
	"learn/helpers" // an imported custom package (module)
)

func main() {
	helpers.Greet("Kenny")
}
