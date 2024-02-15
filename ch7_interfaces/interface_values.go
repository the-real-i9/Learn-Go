package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

/*
A value of an interface type, (or interface value) has two components,
the concrete type: also known as its "dynamic type"
and a value of that type: also known as its "dynamic value".

The "zero value for an interface type" has both is dynamic type and dynamic value set to "nil"

Interface values are comparable using == and !=.
Two interfaces are equal if both have comparable types
AND
if both have each of their components equal to the other.

*/

func intfValues() {
	var w io.Writer // "nil" interface value. dynamic type: nil | dynamic value: nil

	// dynamic type: *os.File |
	// dynamic value: a pointer to os.File variable representing stdout [fd int = 1(stdout)]
	w = os.Stdout
	w.Write([]byte("hello"))

	// to report the dynamic type of an interface value
	fmt.Printf("\n%T\n", w) // *os.File

	// dynamic type: *bytes.Buffer
	// dynamic value: a pointer to the new empty buffer, whose struct fields are initialized to their zero values
	w = &bytes.Buffer{}
	w.Write([]byte("hello"))
	fmt.Println(w)

	w = nil

	// dynamic type: time.Time
	// dynamic value: a time.Time struct variable, whose fields are initialized to their default values
	var x any = time.Now()
	fmt.Println(x)

	intfValues2()
}

/* A nil interface value, which contains no value at all, is not the same as an interface value containing a pointer that happens to be nil */
const debug = true

func intfValues2() {
	// var buf *bytes.Buffer
	var buf io.Writer // fixed

	if debug {
		buf = &bytes.Buffer{}
	}

	// the `out` interface value in f(out io.Writer) can never be "nil" even if debug is false,
	// because the "nil" buf has a dynamic type of *bytes.Buffer and a dynamic value of "nil" pointer.
	// writing to a "nil" pointer cuases a panic, as "nil" is not a valid receiver value for *bytes.Buffer.
	// to fixed this, use an actual "nil" io.Writer from the onset
	f(buf)

	if debug {
		// use buf
		// fmt.Println(buf.String())
		fmt.Println(buf)
	}
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
