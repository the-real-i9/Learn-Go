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

	var iterator func(key string)

	seen := map[string]bool{}

	iterator = func(key string) {
		for _, item := range m[key] {
			if _, ok := m[item]; !ok {
				if !seen[item] {
					finalResultArray = append(finalResultArray, item)
					seen[item] = true
				}
			} else {
				iterator(item)
				if !seen[item] {
					finalResultArray = append(finalResultArray, item)
					seen[item] = true
				}
			}
		}
	}

	for _, key := range sortedMapKeys(m) {
		iterator(key)
		if !seen[key] {
			finalResultArray = append(finalResultArray, key)
			seen[key] = true
		}
	}

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
