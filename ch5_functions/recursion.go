package main

import (
	"fmt"
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
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && strings.HasPrefix(a.Val, "https://") {
				links = append(links, a.Val)
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
