package main

import (
	"fmt"
	// "unicode/utf8"
)

func main() {
	s := "✨😎🎯🎈"
	r := []rune(s)
	fmt.Printf("%x\n", s) // force-breaking to ASCII
	fmt.Printf("%x\n", r)
	fmt.Printf("%U\n", r)
	fmt.Println(string(r))
	fmt.Println(string(rune(97)))
}
