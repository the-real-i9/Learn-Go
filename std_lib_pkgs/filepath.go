package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

/*
- The package filepath provides utility functions for manipulating filename paths in a way compatible with the target OS-defined file paths. It uses either foward slashes or backslashes, depending on the OS.

- To process paths such as URLs that always use forward slashes regardless of the OS, see the path package.

- Consult the documentation for further needs, farewell for now.
*/
func filepath_pkg() {
	// return the absolute representation of this path
	abs_path, err := filepath.Abs("./filepath.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(io.Discard, abs_path)

	// return the last element of a path (trailing slashes are removed)
	fmt.Fprintln(io.Discard,
		filepath.Base("a/b/c"),
		filepath.Base("a/b/c/d.go"),
		filepath.Base("https://a/b/c/index.html"))

	// clean a path
	// filepath.Clean()

	// return the path's current directory
	// it cleans the path
	fmt.Fprintln(io.Discard, filepath.Dir("a/b/c/d.go"))

	// return the path's extension
	fmt.Fprintln(io.Discard, filepath.Ext("./filepath.go")) // .go

	// return the names of all file and directory paths matching the glob pattern
	matches, err := filepath.Glob("/home/*/*")
	if err != nil {
		log.Fatal(err)
	}

	mt_bt, err := json.MarshalIndent(matches, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	io.Discard.Write(append(mt_bt, '\n'))

	// report whether a path is absolute
	// filepath.IsAbs()

	// join any number of path strings into a single path
	// filepath.Join()

	// reports whether a path is local (relative) to directory where this code is evaluated
	// the resource doesn't have to exist actually
	fmt.Fprintln(os.Stdout, filepath.IsLocal("./builtins.go"))  // true: "./ means - in the current directory"
	fmt.Fprintln(os.Stdout, filepath.IsLocal("/builtins.go"))   // false: absolute
	fmt.Fprintln(os.Stdout, filepath.IsLocal("../builtins.go")) // false: "../ means - out of the current directory"
}
