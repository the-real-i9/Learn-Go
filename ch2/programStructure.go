package main

import (
	"fmt"
)

/* -------- */
/* Pointers */
/* -------- */
func pointers() {
	var x int = 5 // x is a variable

	/*
		px is a pointer to x, because it holds its address
		&x is read "address of" x.
	*/
	var px = &x // OR var px *int = &x -- explicit type

	/*
		*px is read "contents" of x
		it refers to the actual variable, x (that px points to)
		and therefore can be reassigned, i.e. we changed the value in x
	*/
	// fmt.Println(*px)
	*px *= 2
	// fmt.Printf("x is %d and *px is %d", x, *px)

	/* var array = []int{1, 2, 3}
	var parr = &array[1]

	fmt.Println(*parr) */

	// var pn *int // fmt.Println(pn == nil) // true

	/* var prf = f()
	*prf *= 5
	fmt.Println(f() == f()) // false */

	/* v := 1
	incr(&v)
	incr(&v)
	fmt.Println(v) */

	/* var a = 22

	var pa = &a

	var pac = pa

	fmt.Println(*pac) */

	/* ----- Pointers: flag ----- */

	/* flag.Parse()
	// `*sep` is what follows `-s` flag in the cmd-line or its default value
	fmt.Print(strings.Join(flag.Args(), *sep))

	// if you don't specify `-n` (no-new-line) in the cmd-line, a new line will print
	if !*noNewLine {
		fmt.Println()
	} */
}

/* // `n` allows us to use the flag `-s` of the cmd-line in our program as a string
var noNewLine = flag.Bool("n", false, "omit trailing newline")

// `sep` allows us to use the flag `-s` of the cmd-line in our program as a string
var sep = flag.String("s", " ", "separator") */

/* func incr(p *int) {
	*p++
} */

/* func f() *int {
	v := 5
	return &v
} */

/* func main() {
	fmt.Print(fib(10))
}

func fib(n int) []int {
	fibSeq := []int{0, 1}
	// x, y := 0, 1

	for i := 0; i < n; i++ {
		// x, y = y, x+y
		fibSeq = append(fibSeq, fibSeq[len(fibSeq)-2]+fibSeq[len(fibSeq)-1])
	}

	return fibSeq
}
*/

/* -------- */
/* Type declarations */
/* -------- */

/*
// declare the type
type Celcius float64
type Farenheit float64

// declare type methods
func (c Celcius) toString() string {
	return fmt.Sprintf("%g^C", c)
}

func (f Farenheit) toString() string {
	return fmt.Sprintf("%g^F", f)
}

// use the type
const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func CToF(c Celcius) Farenheit { return Farenheit(c*9/5 + 32) }
func FToC(f Farenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

func main() {
	fmt.Println(CToF(BoilingC).toString())
} */

/* var x int

func init() {
	x = 5 * 4
}

func main() {
	var y = x
	fmt.Println(y)
} */

/* Scopes */
func f() string { return "Hi" }

var g = "g"

func main() {
	// short declarations shadow (re-declare), outer-scope variables with the same name.
	f := "f"
	fmt.Println(f)
	fmt.Println(g)
}
