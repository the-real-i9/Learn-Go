package main

import "fmt"

func maps() {
	// Map type --- map[key_type]value_type

	// Initializing a Map: Using "make"
	ages := make(map[string]int)
	ages["alice"] = 2
	ages["prince"] = 10

	// Initializing a Map: Using "map literal"
	ages2 := map[string]int{
		"alice": 2,
		"mike":  3,
	}
	// accessing a map
	fmt.Println(ages["alice"])

	// accessing a map actually returns a multi-valued result
	// the first is the result or the zero value of the maps' value type
	// the second is the result of an existential test; sometimes we might only want to test if a key exists (analogous to an `has()` function) or in situations where a map's value is of type integer, and accessing a non-existent key makes its existence ambiguos as we can't tell whether the key really doesn't exist of the value is explicitly "0".

	val, found := ages["mike"]
	fmt.Println(val, found)

	// deleting a key
	delete(ages2, "alice")

	// empty map
	emptyMap := map[string]int{}
	// Zero value for a map type is nil
	fmt.Println(emptyMap)
	// Accessing a non-existent key is the zero value of the map's value type
	fmt.Println(emptyMap["who?"])

	// A map type is enumerable
	for key, value := range ages {
		fmt.Printf("%v's age is %v\n", key, value)
		// NOTE that: the fields of a map are randomly ordered in memory at runtime
		// So if you want an ordered iteration, make sure to
		// sort the keys first and iterate it by accessing with the ordered keys
	}

	// Go does not provide a `set` type, but since the keys of a map are distinct, a map can serve this purpose.
	// A map of type map[string]bool is, in a way, considered a "set" of strings, if the bool is true, though
	setOfStrings := map[string]bool{
		"benny": true,
		"bruno": true,
	}
	// we can do a direct existential test like this, as we we expect true of false
	if setOfStrings["benny"] {
	} // true
	if setOfStrings["ulon"] {
	} // false

	// The key of a map type must be a comparable type. If you want to use a non-comparable type, you must first cast it to a comparable type.
	// Recall how JavaScript's Map can take any type as key.
	sliceKeyMap := map[string]string{}
	sliceKeyMap[fmt.Sprint([]int{2, 4, 5})] = "slice 2,4,5"
}
