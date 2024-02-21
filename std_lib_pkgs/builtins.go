package main

import (
	"fmt"
)

func builtins() {
	/* =========== */
	/* slc := []int{1, 2, 3, 4, 5}
	mp := map[string]byte{"a": 'a', "b": 'b'}

	clear(slc)
	clear(mp)

	fmt.Println(slc)
	fmt.Println(mp) */

	/* ===== copy() ====== */
	/* src := []byte("kenny and tenny")
	dst := make([]byte, len(src), cap(src))
	// src_alt := "kenny and tenny"
	// dst_alt := make([]byte, len(src))

	// dst cannot be a nil slice for copy
	copy(dst, src)

	fmt.Println(string(dst)) */

	/* ===== append() ====== */
	/* byte_str := []byte("kenny and tenny")
	byte_dst := []byte{}

	byte_dst = append(byte_dst, byte_str...)

	fmt.Println(string(byte_dst)) */

	/* ===== max(), min() ====== */
	/* nums := []int{1, 2, 3, 4, 5, 6}
	alphas := []string{"a", "b", "c", "d"}

	// stil confused
	fmt.Println(max(alphas[0], alphas[1])) */

	/* =========== */
	/* ptr_int := new(int)        // *ptr_int == 0
	ptr_str := new(string)     // *ptr_int == ""
	ptr_slice := new([]string) // *ptr_slice == nil or []

	*ptr_int = 5
	*ptr_str = "kenny"
	*ptr_slice = append(*ptr_slice, "ken", "sam")

	fmt.Println(*ptr_int)
	fmt.Println(*ptr_str)
	fmt.Println(*ptr_slice) */

	/* ====== panic(), recover() ===== */
	/* defer func() {
		// panic + recover (in a deffered function in the panicking function)
		// will intercept(stop) the panicking sequence, retrieve its error value, handle it,
		// and restore normal execution
		// causing the panicking function to return normally (and immediately)
		p := recover()
		fmt.Printf("panic: %v\n", p)
	}()

	if true {
		// panic alone will propagate up to the
		// main goroutine and terminate the program
		panic("What the heck was that???!")
	} */

	/* ==== others ==== */
	byte_is := []byte("sam")
	is_uint8 := []uint8("sam")

	rune_is := []rune("ken")
	is_int32 := []int32("ken")

	fmt.Println(string(byte_is) == string(is_uint8))
	fmt.Println(string(rune_is) == string(is_int32))
}
