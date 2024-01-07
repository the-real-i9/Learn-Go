package main

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

var P1 = Person{name: "Kenny", age: 23, Sibling: &Sibling{Person: Person{name: "Amos", age: 26, Sibling: &Sibling{Person: Person{name: "Tosin", age: 28}, Role: "Elder sister"}}, Role: "Elder brother"}}

/* ------------- */

func Visit(links []string /* , n *html.Node */) []string {
	// check if this node is a accessible html tag with an hyperlink e.g. a, img, audio, video etc. If so extract its link attribute (href or source) and add it to the links slice.
	// check if this node has a valid Sibling, if yes, call this function on this node again, passing it our links slice.
	// This function recursively visits nested nodes to grab all the links it finds.

	return []string{""}
}
