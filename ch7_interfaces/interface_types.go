package main

import (
	"fmt"
	"io"
)

/*
An interface type specifies a set of methods that
a concrete type must possess to be considered
an instance of that interface
*/

type StringReader string

func (mr StringReader) Read(p []byte) (int, error) {
	copy(p, mr)
	return len(mr), io.EOF
}

func NewReader(s string) io.Reader {
	var mr = StringReader(s)
	return mr
}

type LimitedReader struct {
	R io.Reader
	N int64
}

func (l LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}

	return l.R.Read(p)
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return LimitedReader{R: r, N: n}
}

func intfTypes() {
	rd := NewReader("kenny is a boy")

	lrd := LimitReader(rd, 10)

	bts, err := io.ReadAll(lrd)
	if err == io.EOF {
		fmt.Println("err is EOF")
	}

	fmt.Println(string(bts))
}
