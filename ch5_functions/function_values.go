package main

import (
	"fmt"
	"strings"
)

// Function are first-class values; like other values, they have types, they may be assigned to variables, passed to or returned from functions. A function value may be called like any other function.

// functions values are represented by the function type followed by the function body

// what's before the open curly brace is technically a complete function type
func square(n int) int {
	return n * n
}
func doSomething(n int) {
	/* doing something */
}

// the name preceding the parenthenses is technically the name of the variable the function is being assigned to, for a function declaration

// consider an explicit function value assignment to a variable
// implicit typing
var negative = func(n int) int {
	return -n
}

// explicit typing
var product func(m, n int) int = func(m, n int) int {
	return m * n
}

// Uninitialized function i.e. a function with no value/
// The zero value of a function type is nil. They can be compared with nil. But they are not comparable
var uf func(int) int

// A comlete function type includes the "func" keyword, the number function arguments, and the number function results.
// In explicit typing, the function value you assign must match the type defined.

func funcVals() {
	uf = square
	uf(5)

	sq := square(5)
	neg := negative(5)
	prod := product(5, 6)
	doSomething(5)

	fmt.Println(sq, neg, prod)

	// functions let us parameterize our functions over not just data, but behavious too. That is, function values can be passed as arguments to functions.
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "kehinde"))
}
