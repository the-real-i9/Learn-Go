package main

import (
	"fmt"
)

func sum(vals ...int) int {
	// a slice is created for the "vals" argument and the arguments are packed into it
	// analogous to the javascript rest operator
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}

func max(firstVal int, vals ...int) int {
	max := firstVal

	for _, val := range vals {
		if val > max {
			max = val
		}
	}

	return max
}

func min(firstVal int, vals ...int) int {
	min := firstVal

	for _, val := range vals {
		if val < min {
			min = val
		}
	}

	return min
}

func variadicFuncs() {
	// fmt.Println(sum(1, 2, 3, 4, 5))
	// fmt.Println(sum([]int{6, 7, 8, 9, 10}...)) // analogous to the javascript spread operator
	fmt.Println(max(3, 5, 2, 2, 5))
	fmt.Println(min(3, 5, 2, 2, 5))
}
