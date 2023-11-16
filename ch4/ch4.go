package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/* Array
func main() {
	var arr3 = [4]int{1, 2, 3, 4}

	var arr = [...]int{1, 2, 3, 4}
	var arr2 = [...]int{1, 2, 3, 4}
	fmt.Printf("%t", arr == arr2)

	var arr3 = [3]int{1, 2, 3, 4} // out of bounds
} */

/* type Currency int

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
*/

/* // Struct

type Staff struct {
	ID                 string
	Firstname, surname string
}

type Manager struct {
	Staff // Staff
	Role string
}

func main() {
	// var Person1 Staff = Staff{ID: "fkjk3j4", Firstname: "Kehinde", surname: "Ogunrinola"}

	// var Person2 Staff = Staff{"maelj2l", "Samuel", "Oluwarinola"}

	// we can create a pointer to struct, just like a normal variable
	// Person3.ID --- dot-notation works directly on pointer to struct
	// var Person3 *Staff = &Staff{ID: "ki422"}

	var Person4 Manager = Manager{Staff: Staff{Firstname: "Dan"}, Role: "Manager"}

	fmt.Printf("%v", Person4)
} */

// JSON
type Movie struct {
	Title  string
	Year   int  `json:"released"`        // "released" as key
	Color  bool `json:"color,omitempty"` // "color" as key, omit if empty
	Actors []string
	// the left-trailing comma is so as not to make the compiler think we want "omitempty" as key
	TimesWatched int    `json:",omitempty"` // struct Field as key, omit if empty
	DevComment   string `json:"-"`          // omit
	// they right-trailing comma is so as not to make the compiler thing we want to omit
	HyphenKey bool `json:"-,"` // you want "-" as key
}

// empty detection, works based on JavaScript's falsy values

var movies = []Movie{{Title: "Casablanca", Year: 1942, TimesWatched: 0, Actors: []string{"Humphrey Bogart", "Ingird Bergman"}, DevComment: "testing"}}

// var moviesJSONDecoded []Movie
var moviesJSONDecoded = []Movie{}

func main() {
	// data, _ := json.Marshal(movies)
	dataIndent, _ := json.MarshalIndent(movies, "", "  ")
	// fmt.Printf("%s\n", dataIndent)

	// fmt.Printf("%v\n", moviesJSONDecoded) // []

	json.Unmarshal(dataIndent, &moviesJSONDecoded)
	// fmt.Printf("%v\n", moviesJSONDecoded)

	envJson, _ := json.MarshalIndent(os.Environ(), "", "  ")
	fmt.Printf("%s\n", envJson)
}
