package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func typeAssert() {
	var w io.Writer = os.Stdout

	/* If the asserted type T is a concrete type, then the type asserted checks whether the interface's dynamic type is identical to T.
	 */
	wt := w.(*os.File) // correct: w's dynamic type is *os.File
	// c := w.(*bytes.Buffer) // wrong: w's dynamic type is not *bytes.Buffer

	/* A type assertion to a concrete type extracts the concrete value from its operand */
	fmt.Fprintln(wt, "ken") // wt == os.Stdout

	// typeAssert2()
	// typeAssert3()
	// typeAssert4()
	typeAssert5()
}

func typeAssert2() {

	/* If the type asserted type T is an interface type, then the type assertion checks whether the interface's dynamic type satisfies T */
	var w io.Writer = os.Stdout // w's concrete type is *os.File but holds only the Write() method

	/* A type assertion to an interface type changes the type of the expression, making a different set of methods accessible, but it preserves the dynamic type and value components inside the interface value */
	/* It rightly assumes that since you're asserting a type you probably need the methods of that type, even if that wasn't the original one it was assigned to. */
	rw := w.(io.ReadWriter) // rw's concrete type is also *os.File but holds both the Read() and Write() method

	fmt.Fprintf(rw, "%T\n", rw)
}

func typeAssert3() {
	/* If the type assertion in an assignment in which two results are expected, the operation does not panic on failure but instead returns an additional second result, a boolean indicating true/false */

	var w io.Writer = os.Stdout
	_, fok := w.(*os.File)
	fmt.Println(fok) // true

	_, bok := w.(*bytes.Buffer)
	fmt.Println(bok) // false

	_, rwok := w.(io.ReadWriter)
	fmt.Println(rwok)

	/* Common usage */
	if w, ok := w.(*os.File); ok {
		// using w
		fmt.Fprintln(w, "Yo men!")
	}

	w = os.Stdin
	// berofe assertion, w only exposes Write()
	if w, ok := w.(io.ReadWriteCloser); ok { // ok == true, since os.Stdin satisfies the io.ReadWriteCloser
		// using `w`-- after assertion, `w` now exposes Read() and Close()
		scanner := bufio.NewScanner(w)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
			w.Close()
		}
	}
}

/* Discriminating Errors with Type Assertions:
This is used in situations where an error returned by an attempted file operation
may be one of several kinds of errors in which you need to handle each differently */

func typeAssert4() {
	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err)) // deprecated: It only supports errors returned by the `os` package

	// new
	errors.Is(err, fs.ErrNotExist /* file does not exist */)
	errors.Is(err, fs.ErrExist /* file already exists */)
	errors.Is(err, fs.ErrPermission /* permission denied */)
	errors.Is(err, fs.ErrClosed /* file already closed */)
}

/* Type Switches: Control flow based on the asserted type of an empty inteface{} */
func typeAssert5() {
	fmt.Println(whatType(2 == 5))
}

func whatType(x interface{}) string {
	// switch x.(type) { // in case you don't need the value
	switch x := x.(type) {
	case nil:
		return fmt.Sprintf("%v is %q", x, "nil")
	case int:
		return fmt.Sprintf("%v is an %q", x, "integer")
	case bool:
		if x {
			return fmt.Sprintf("%v is a %q", x, "true boolean")
		}
		return fmt.Sprintf("%v is a %q", x, "false boolean")
	default:
		panic("Stop that!")
	}
}
