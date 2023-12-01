package main

import "fmt"

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

func main() {
	fmt.Println(p1.sibling)
}
