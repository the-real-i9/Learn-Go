package main

import (
	"fmt"
	"os"
)

func fmt_pkg() {
	printers()
	// scanners()
}

func printers() {
	/*
		* Format suffixes: They apply before the output prefixes
			* No-suffix: Appends no newline to the string
			* ln-suffix: Appends newline to the string. Other suffixes should append \n to the string, if newline is desired
			* f-suffix: Generates a result string from a format string, followed by format specifier arguments

		* Output prefixes: They apply after the format suffixes
			* No-prefix: Prints the (resulting) string to the standard output stream
			* F-prefix: Writes the (resulting) string to the instance of io.Writer passed to it
			* S-prefix: Returns the (resulting) string
	*/

	fmt.Print("Good morning\n") // newline added manually, if desired
	fmt.Fprint(os.Stdout, "Good afternoon\n")
	str := fmt.Sprint("Good evening\n")

	os.Stdout.WriteString(str)

	fmt.Println("How's everything?")
	fmt.Fprintln(os.Stdout, "All's well, I guess.")
	str2 := fmt.Sprintln("See you later.")

	os.Stdout.WriteString(str2)

	fmt.Printf("How's %s\n", "your day?") // newline added manually, if desired
	fmt.Fprintf(os.Stdout, "How's %s\n", "work?")
	str3 := fmt.Sprintf("How's %s\n", "your night?")

	os.Stdout.WriteString(str3)

}

func scanners() {
	/*
		* Input prefixes: They apply before the format suffixes
			* No-prefix: Scans text from the standard input stream
			* F-prefix: Scans text from the instance of io.Reader passed to it
			* S-prefix: Scans text from the string from the one passed to it

		* Format suffixes: They apply after the input prefixes
			* No-suffix: Stores successive space separated values into sucessive arguments. Newlines are regarded as spaces until there all arguments are supplied values.
			* ln-suffix: Like "No-suffix" but stops scanning at a newline.
			* f-suffix: Stores successive space-separated values into successive arguments as determined by the format.
				* Newlines in the input must match newlines in the format, if not it stops scanning at a newline. Except if the format specifier being supplied is %c, in which case it will accept any rune, including newlines.
	*/

	var (
		str      string
		str_next string
	)

	fmt.Println("---- No newline scan ----")

	fmt.Scan(&str, &str_next)
	fmt.Println(str, str_next)

	fmt.Fscan(os.Stdin, &str, &str_next)
	fmt.Println(str, str_next)

	fmt.Sscan("Hello world", &str, &str_next)
	fmt.Println(str, str_next)

	var (
		str_ln      string
		str_next_ln string
	)

	fmt.Println("---- Newline scan ----")

	fmt.Scanln(&str_ln, &str_next)
	fmt.Println(str_ln, str_next_ln)

	fmt.Fscanln(os.Stdin, &str_ln, &str_next_ln)
	fmt.Println(str_ln, str_next_ln)

	fmt.Sscanln("Hello world", &str_ln, &str_next_ln)
	fmt.Println(str_ln, str_next_ln)

	var (
		product  string
		price    float64
		currency string
	)

	fmt.Println("---- Format scan ----")

	fmt.Scanf("%s is sold for %g %s", &product, &price, &currency)
	fmt.Printf("product: %s, price: %.1f, currency: %s\n", product, price, currency)

	fmt.Fscanf(os.Stdin, "%s is sold for %g %s", &product, &price, &currency)
	fmt.Printf("product: %s, price: %.1f, currency: %s\n", product, price, currency)

	fmt.Sscanf("Gold is sold for 10000 karat", "%s is sold for %g %s", &product, &price, &currency)
	fmt.Printf("product: %s, price: %.1f, currency: %s\n", product, price, currency)
}
