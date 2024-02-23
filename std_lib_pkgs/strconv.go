package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

/* Think about `strconv` when you think of any string or number conversion */
func strconv_pkg() {
	// formatters()
	// parsers()
	// quoters()
	appenders()
	// others()
}

func formatters() {
	fmt.Println(strconv.FormatInt(10, 16)) // a
	fmt.Println(strconv.FormatUint(10, 2)) // 1010
	fmt.Println(strconv.FormatFloat(3.142, 'g', -1, 64))
	// you'll use the 'g' or 'G' for `fmt` in most cases,
	// as it automatically decides to use the 'e' standard form (or not) in ideal cases
	// others `fmt`'s are: 'b' for binary float and 'x' for hexadecimal float

	fmt.Println(strconv.FormatBool(true)) // true

	/* Note that: the `fmt` package format variants can do most of the strconv.Format... 's job */
}

func parsers() {
	// discovery: `int` type is 64bit integer in representation, like int64
	i, err := strconv.ParseInt(fmt.Sprint(math.MaxInt), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)

	i2, err := strconv.ParseInt("1101001010010101", 2, 0) // variant: ParseUint()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i2)

	ui, err := strconv.ParseUint("110010011", 2, 8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ui)

	/* === others ===== */
	// strconv.ParseFloat("3.14", 64)
	// strconv.ParseBool("true")
}

func quoters() {
	fmt.Println(strconv.Quote("mess"))                     // "mess"
	fmt.Println(strconv.QuoteRune('💔'))                    // '💔'
	fmt.Println(strconv.QuoteToASCII("I ❤ you!	"))         // "I \u2764 you\t"
	fmt.Println(strconv.QuoteRuneToASCII('	'))             // '\t'
	fmt.Println(strconv.QuoteRuneToASCII('❤'))             // '\u2764'
	fmt.Println(strconv.QuoteToGraphic("I \u2764 you!\t")) // "I ❤ you!\t"
	fmt.Println(strconv.QuoteRuneToGraphic('\u2766'))      // '❦'
}

/*
Appenders are like formatters and quoters, coupled with
the feature of appending the result to an existing byte slice
*/
func appenders() {
	love := []byte("love: ")
	love = strconv.AppendQuoteRune(love, '❤')
	love = strconv.AppendQuote(love, "❤ you")
	fmt.Println(string(love))
}

func others() {
	// common
	fmt.Println(strconv.Atoi("56543"))
	fmt.Println(strconv.Itoa(56543))
}
