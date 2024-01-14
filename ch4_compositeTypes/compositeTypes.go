package main

import (
	"fmt"
	"log"
)

func main() {
	result, err := SearchGithubIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	// maps()
}
