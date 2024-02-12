package main

import "fmt"

// named "struct" type
type Staff struct {
	ID                 string
	Firstname, surname string /* Exported, unexported */
}

type Manager struct {
	Staff // Staff
	Role  string
}

// A named struct cannot contain a field of its own type, execept we make it a pointer
// This makes recursive data structures possible, trees for example.
type node struct {
	data        string
	firstChild  *node
	nextSibling *node
	chilren     []node
}

// The zero value of a struct type is the zero value of all its fields' type
var myNode = node{} // an empty struct

func structType() {
	var Person1 Staff = Staff{
		ID:        "fkjk3j4",
		Firstname: "Kehinde",
		surname:   "Ogunrinola",
	}

	// if you don't specify the keys, then you must arrange the values in the same order in which it is declared
	var Person2 Staff = Staff{"maelj2l", "Samuel", "Oluwarinola"}

	// we can create a pointer to struct, just like a normal variable
	var Person3 *Staff = &Staff{ID: "ki422"}

	var Person4 Manager = Manager{Staff: Staff{Firstname: "Dan"}, Role: "Manager"}

	Person4.Firstname = "Cole"

	// the dot-notation works on a normal struct or a pointer to struct
	fmt.Printf("%v", Person4.Firstname)
	fmt.Printf("%v", Person3)
	fmt.Printf("%v", Person2.Firstname)
	fmt.Printf("%v", Person1.surname)

	// A struct is comparable if all the fields of the struct are comparable
	// In which case we can use it as key in a map
}

// If you want a named struct function-result be writable, then make it addressable,  else the result is read-only.
// This equally applies to struct arguments, if you want them to be writable within the function
// This means struct are call-by-value types
func structReturnWritable(staff *Staff, newSurname string) *Staff {
	staff.surname = newSurname // writable
	return staff               // staff is already a pointer
	// return &staff					 // if staff is not already a pointer
}

func structReturnReadOnly(staff Staff, newSurname string) Staff {
	staff.surname = newSurname
	return staff
}

// unnamed "struct" type
// cannot have methods, unlike named types
// this  is just like declaring a mere (unnamed) varable of any type
var student struct {
	name  string
	class string
} = struct {
	name  string
	class string
}{"taye", "B"}

var teacher = struct {
	name  string
	class string
}{"taye", "B"}
