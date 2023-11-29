package main

import (
	"fmt"
)

/* func main1() {
	var arr3 = [4]int{1, 2, 3, 4}

	var arr = [...]int{1, 2, 3, 4}
	var arr2 = [...]int{1, 2, 3, 4}
	fmt.Printf("%t", arr == arr2)

	var arr3 = [3]int{1, 2, 3, 4} // out of bounds
} */

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

var symbol = [...]string{USD: "USD", EUR: "EUR"}

func main() {
	runes := []rune("Hello ₶₹₰₫₦")

	fmt.Printf("%q ", runes)
}
