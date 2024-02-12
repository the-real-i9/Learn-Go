package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// a concrete type ByteCounter that implements the io.Writer interface
type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (int, error) {
	*bc += ByteCounter(len(p))

	return len(p), nil
}

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	count := len(bytes.Split(p, []byte(" ")))
	*wc += WordCounter(count)

	return count, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	count := len(bytes.Split(p, []byte("\n")))

	*lc += LineCounter(count)

	return count, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	content := []byte("brosky")
	count := int64(len(content))

	w.Write(content)

	return w, &count
}

var bc ByteCounter
var wc WordCounter
var lc LineCounter

func intfAsCont() {
	fmt.Fprintf(&bc, "hello, %s", "Kenny")
	fmt.Println(bc)

	io.Copy(&wc, os.Stdin)
	fmt.Fprintf(&wc, "taye is a boy and a girl")
	fmt.Println(wc)

	io.Copy(&lc, os.Stdin)
	fmt.Fprintf(&lc, "taye \n is a boy \n and a girl")
	fmt.Println(lc)

	_, n := CountingWriter(os.Stdout)
	fmt.Printf("\n%d\n", *n)

}
