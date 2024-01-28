package main

import (
	"fmt"
	"slices"
)

/*
- These are: Function values used within an expression,
rather than at the package level.
- They are denoted by function literals, i.e. functions that
do not have names in contrast to function declarations.
- Functions defined in this way have access to the entire
lexical environment, so the inner function can refer to
variables from the enclosing function.
*/

/*
The anonymous inner function can access and update the local
variables of the enclosing function. These hidden variable
references are why we classify functions as reference
types. Function values like these are implemented using a
technique called closures.

Recall that, the lifetime of a variable
is not determined by its scope.
*/
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func sortedMapKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func topoSort(m map[string][]string) []string {
	finalResultArray := []string{}

	var visitAll func(items []string)

	seen := map[string]bool{}

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				finalResultArray = append(finalResultArray, item)
			}
		}
	}

	visitAll(sortedMapKeys(m))

	return finalResultArray
}

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func anonymFuncs() {
	/* f := squares()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f()) */

	for _, course := range topoSort(prereqs) {
		fmt.Println(course)
	}
}

/* NOTES */

/* WARNING!!! Anonymous functions created in an iteration generally does capture the loop's iteration variable -
its address, to be precise - therefore, the function's operations on this iteration variable are performed
on the same variable (address) during each iteration. To avoid this behaviour, declare a new variable
in each iteration and assign and copy the iteration variable into it via assignment.
*/
