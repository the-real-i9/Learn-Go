package main

import "fmt"

type Staff struct {
	ID                 string
	Firstname, surname string /* exported, unexported */
}

type Manager struct {
	Staff // Staff
	Role  string
}

func Struct() {
	var Person1 Staff = Staff{ID: "fkjk3j4", Firstname: "Kehinde", surname: "Ogunrinola"}

	var Person2 Staff = Staff{"maelj2l", "Samuel", "Oluwarinola"}

	// we can create a pointer to struct, just like a normal variable
	// Person3.ID --- dot-notation works directly on pointer to struct
	var Person3 *Staff = &Staff{ID: "ki422"}

	var Person4 Manager = Manager{Staff: Staff{Firstname: "Dan"}, Role: "Manager"}

	fmt.Printf("%v", Person4)
	fmt.Printf("%v", Person3)
	fmt.Printf("%v", Person2)
	fmt.Printf("%v", Person1)
}
