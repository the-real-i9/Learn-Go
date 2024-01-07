package main

import (
	"fmt"
	"log"
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
	result, err := SearchGithubIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
