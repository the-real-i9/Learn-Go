package main

import (
	"fmt"
	"index/suffixarray"
	"os"
	"regexp"
)

/* The index/suffixarray implements substring search in logarithmic time using an in-memory suffix array
This is in contrast to regexp, which using linear time.

Kinda like database "indexing".
*/

func index_pkg() {
	str := "file: go.mod, file: main.c"
	// creating an index
	index := suffixarray.New([]byte(str))

	// find all indexes where the specified byte slice occured
	fmt.Println(index.Lookup([]byte("file"), -1 /* number of indexes to return */))

	// to lookup with regexp instead, use FindAllIndex()
	// works like (*regexp.Regexp).FindAllIndex(), but in logarithic time sa described above

	// Recall that: The index is an integer pair in []int, which actually defines the lower and upper boundaries of a slice operation, able to extract the substring from the original string
	indexes := index.FindAllIndex(regexp.MustCompile(`\w+(\.[a-z]+)`), -1)
	fmt.Println(indexes)

	// using the found indexes to extract the target substrings in constant time
	strs := make([]string, len(indexes))
	for i, pair := range indexes {
		strs[i] = str[pair[0]:pair[1]]
	}

	// it should print valid file names with extensions
	fmt.Println(strs) // [go.mod main.c]

	/* Other methods */
	// index.Read(io.Reader) // reads the index from io.Reader into index
	// index.Write(io.Writer) // writes the index to io.Writer

	// str2 := "file: main.cpp, file: main.go"
	// reader := strings.NewReader(str2)
	if err := index.Read(os.Stdin); err == nil {
		index.Write(os.Stdout)
		fmt.Println()
	} else {
		fmt.Println(err)
	}
}
