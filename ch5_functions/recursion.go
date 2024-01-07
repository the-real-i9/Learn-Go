package main

import (
	"fmt"
)

// You can use this recursion to mimic prototypal inheritance like in JavaScript
type Person struct {
	name string
	age  int
	/*
	* Normally, sibling is a type of Person, but its a person with a role, which is, of course, a type of Sibling.
	* I love this thinking: Every Person can have a Sibling, and every Sibling can as well have a Sibling since a Sibling is a Person
	 */
	sibling *Sibling
	parent  *Parent
}

type Sibling struct {
	Person
	role string
}

type Parent struct {
	Person
	role string
}

var p1 = Person{name: "Kenny", age: 23, sibling: &Sibling{Person: Person{name: "Amos", age: 26, sibling: &Sibling{Person: Person{name: "Tosin", age: 28}, role: "Elder sister"}}, role: "Elder brother"}}

func main1() {
	fmt.Println(p1.sibling)
}

/* ------------- */

func visit(links []string /* , n *html.Node */) []string {
	// check if this node is a accessible html tag with an hyperlink e.g. a, img, audio, video etc. If so extract its link attribute (href or source) and add it to the links slice.
	// check if this node has a valid sibling, if yes, call this function on this node again, passing it our links slice.
	// This function recursively visits nested nodes to grab all the links it finds.

	return []string{""}
}

func main() {

}
