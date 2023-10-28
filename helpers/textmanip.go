package helpers

// this is a package
// note that it's same with a module

import "fmt"

// note the package func name begins with uppercase, this is a must

// a "func name starting with lowercase is not a package func", but a normal func that merely groups related statements, "to be used" in the file or another file in the module/package (as an exported func). In fact, not using it will throw error

// in other packages, including "main", after importing this package (import "learn/helpers"), you can access it with "helpers.Greet()"
func Greet(name string) {
	fmt.Printf("Hey, %s!", name)
}
