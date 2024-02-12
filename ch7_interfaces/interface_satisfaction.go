package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func intfSatsf() {
	// a concrete type satisfies an interface if it implements all of its methods
	var w io.Writer

	/* even though the concrete types have several methods of their own,
	the io.Writer only exposes its own behaviour Write(), wrapped and concealed */
	w = os.Stdout

	/* (*bytes.Buffer) implements the Write() method and therefore satisfies the interface io.Writer,
	hence, we assign (*T) rather than (T) */
	w = &bytes.Buffer{}

	w.Write([]byte("hello"))
	fmt.Println(w)

	// an empty interface is satisfied by any concrete type
	var x any // type (any) is same as empty interface (interface{})
	x = true
	x = 12.34
	x = "hello"
	x = map[string]int{"one": 1}
	x = bytes.Buffer{}
	fmt.Println(x)
}
