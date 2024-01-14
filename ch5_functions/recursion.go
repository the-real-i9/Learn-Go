package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

// You can use this recursion to mimic prototypal inheritance like in JavaScript
type Person struct {
	name string
	age  int
	/*
	* Normally, Sibling is a type of Person, but its a person with a Role, which is, of course, a type of Sibling.
	* I love this thinking: Every Person can have a Sibling, and every Sibling can as well have a Sibling since a Sibling is a Person
	 */
	Sibling *Sibling
	Parent  *Parent
}

type Sibling struct {
	Person
	Role string
}

type Parent struct {
	Person
	Role string
}

var Person_1 = Person{
	name: "Kenny",
	age:  23,
	Sibling: &Sibling{
		Person: Person{
			name: "Amos",
			age:  26,
			Sibling: &Sibling{
				Person: Person{
					name: "Tosin",
					age:  28,
				},
				Role: "Elder sister",
			},
		},
		Role: "Elder brother",
	},
}

/* ------------- */
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && slices.Contains([]string{"a", "img", "script", "link", "video", "audiio"}, n.Data) {
		for _, attr := range n.Attr {
			if slices.Contains([]string{"href", "src"}, attr.Key) && strings.HasPrefix(attr.Val, "https:") {
				links = append(links, fmt.Sprintf("%v --- %v", n.Data, attr.Val))
			}
		}
	}

	/* for c := n.FirstChild; c != nil; c = c.NextSibling {
	links = visit(links, c)
	} */
	c := n.FirstChild
	if c != nil {
		links = visit(links, c)
	}

	next := n.NextSibling
	if next != nil {
		links = visit(links, next)
	}

	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func findlinks1(webPage io.Reader) {
	doc, err := html.Parse(webPage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func findtags(webPage io.Reader) {
	doc, err := html.Parse(webPage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findtags1: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc)
}

func trackTags(elCounts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elCounts[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elCounts = trackTags(elCounts, c)
	}

	return elCounts
}

func countElements(webPage io.Reader) {
	doc, err := html.Parse(webPage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countElements: %v\n", err)
		os.Exit(1)
	}

	for k, v := range trackTags(map[string]int{}, doc) {
		fmt.Printf("Tag: %v --- Count: %v\n", k, v)
	}
}
