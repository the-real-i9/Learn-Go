package main

import (
	"log"
	"os"
)

/*
* The log package implements a Logger and functions to manage it.
* The default logger can be used with its predefined settings, but a new one can initialized.

* A logger has three components:
- An io.Writer instance to write the message
- A prefix string that prefixes each message
- A flag for attaching properties like date, time, line number etc.

* A newly created logger is provided with these parameters.
* Both the default logger and a new logger can modified these settings with:
`setOutput()`, `setPrefix()` and `setFlags()` functions and methods, respectively.

* The default logger uses the os.Stderr as its io.Writer instance
*/
func log_pkg() {
	f, err := os.OpenFile("test-error.log", os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// create a logger
	logger := log.New(f, "error: ", log.Lshortfile)

	logger.SetPrefix("errorneous: ")
	logger.SetFlags(log.Ltime)
	logger.SetOutput(os.Stdout)

	/* === Method Variants === */
	/*
		- The fmt.Print suffixes (none, ln and f) have the usual behaviour
		- Generally, all its printer methods work like fmt's Printer, but with the following improvement:

		- For "Panic" variants, printing is followed by a panic
		- For "Fatal" variants, printing is followed by os.Exit(1)
	*/

	logger.Println("Oh no way!")

	// Prefix(), Flags() and Output() lets you retrieve the current set value for the logger's prefix, flags and io.Writer, respectively.
}
