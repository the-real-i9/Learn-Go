package main

import (
	"fmt"
	"slices"
	"strings"
)

/* You're already familiar with most of the functions.
- Just document the review the ones you're most likely to use, including unfamiliar ones

- Method variants with "Func" generally allow you to define your own rules in which element(s) must satisfy to be considered valid
*/

func slices_pkg() {
	/* ==== Compact[Func]() ==== */
	// Compact removes consecutive duplicate elements from a slice
	// Note: The slice must have a comparable element type

	slc_cmpct := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 6, 6, 7, 8, 7, 9}

	fmt.Println(slices.Compact(slc_cmpct))

	// A variant, CompactFunc() allows an equality function to compare element pairs
	// It relaxes the constraint that elements must be comparable, since you dictate the comparison rules
	slc_cmpct_func := []string{"bob", "Bob", "alice", "Vera", "VERA"}
	fmt.Println(slices.CompactFunc(slc_cmpct_func, func(a string, b string) bool {
		// return strings.ToLower(a) == strings.ToLower(b)
		return strings.EqualFold(a, b)
	}))

	fmt.Println("----------------------")

	/* ====== Compare[Func]() ======== */
	// Generally, slices are non-comparable types by themselves, they can only be compared to nil
	// Compare() does an equality test between two slices, returning 0 if s1 == s2, -1 if s1 < s2, and 1 if s1 > s2
	// To avoid explanation, agree that this would work intuitively
	// Note: It requires that s1 and s1 be of the same element types

	slc_cmp_1 := []string{"Ade", "Bade", "Shade"}

	fmt.Println(slices.Compare(slc_cmp_1, []string{"Ade", "Bade", "Shade"})) // 0
	fmt.Println(slices.Compare(slc_cmp_1, []string{"Ade", "Shade", "Bade"})) // -1
	fmt.Println(slices.Compare(slc_cmp_1, []string{"Shade", "Ade", "Bade"})) // -1
	fmt.Println(slices.Compare([]string{"Shade", "Ade", "Bade"}, slc_cmp_1)) // 1
	fmt.Println(slices.Compare([]string{"Shade", "Ade", "Bade"}, slc_cmp_1)) // 1

	// A varaiant, CompareFunc(), allows you define your comparison function
	// Therefore, the types of s1 and s2 don't have to the same since your compare function should handle the differences

	fmt.Println("---------- Concat() ------------")

	/* === Concat() === */
	// Merges multiple slices and return the resulting merged slice

	fmt.Println(slices.Concat([]int{1, 2, 3}, []int{4, 5, 6}))

	fmt.Println("--------- Min(), Max() -------------")

	/* ======= Min[Func](), Max[Func]() ========= */

	fmt.Println(slices.Max([]int{1, 2, 5, 6, 5, 3}))
	fmt.Printf("%q\n", slices.Min([]rune{'1', '2', '5', '6', '5', '3'}))

	// The "Func" variant takes a comparison function that compares element pairs

	fmt.Println("--------- Replace() ---------")

	/* ======= Replace() ======= */
	slc_repl := []string{"Alic", "Bob", "Vera", "Zac"}
	fmt.Println(slices.Replace(slc_repl, 1, 3, []string{"Bill", "Billie", "Cat"}...))

	/* ============================== */
	/* ====== Others/Familiar ======= */
	// Contains[Func]() -- does this slice contain an element with this value? The "Func" variant allows you specify the rule that an element must statisfy to be considered the target value.

	// Delete() -- removes one or more elements from the slice as specified by the boundary (i, j int).

	// DeleteFunc() -- removes one or more elements from the slice for which func() returns true

	// Equal[Func]() -- reports whether two slices are equal. The "Func" variant allows an equality function to define your own equality terms.

	// Index[Func]() -- returns the first occurence of a specified value in a slice. The "Func" variant allows you specify the rule that an element must statisfy to be considered the target value.

	// Insert() -- inserts arbitrary number of elements at the specified index

	// IsSorted[Func]() -- tests whether slice is sorted in ascending order. The "Func" variant allows a comparison function

	// Reverse() -- reverses the elements of a slice in place

	// Sort[[Stable]Func]() -- Sorts the elements of a slice in place. The "Func" varant allows a comparison function for flexibility. The "Stable" variant of "Func" sorts the slice while keeping the original order of equal elements.

}
