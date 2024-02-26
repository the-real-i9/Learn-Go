package main

import (
	"fmt"
	"regexp"
)

func regexp_pkg() {
	// compilers()
	// matchers()
	// finders()
	// replacers()
	regexp_others()
}

func compilers() {
	// to create the regular expression
	// this is the method you're most likely to use
	// the variant is: CompilePOSIX. You'll sparinly need it

	// this regexp matches the extension of a file name
	ext_re := regexp.MustCompile(`(\.[a-z]+)$`) // returns a *Regexp object

	// test it
	fmt.Println(ext_re.FindString("file.data.go"))
}

func matchers() {
	/*
		   Matchers are testers, they're for testing a regexp against a string or []byte
		   - They can be accessed at the package level with a regexp string,
			 - or as the method of a pre-compiled regexp string

			Variants: The variants vary by the type of input
			- MatchString([re string, ]str string) where input is a string,
			- MatchReader([re string, ]r io.RuneReader) where input is a rune from io.RuneReader
	*/

	// Package level
	if matched, err := regexp.Match(`(\.[a-z]+)$`, []byte("go.mod")); err == nil {
		// if matched { ... } else { ... }
		fmt.Println(matched) // true
	}

	// Method of a compiled regexp
	re := regexp.MustCompile(`(\.[a-z]+)$`)
	fmt.Println(re.Match([]byte("go.wor_k"))) // false

}

func finders() {
	/*
		The desired finder method is selected based on this desicion pattern:
		Find(All)?(String)?(Submatch)?(Index)?
		- All : Do you want find "all" the results or the "first"?
		- String : Do you want to work with a "T<string>" or "T<[]byte>" as argument and/or return value
		- Submatch : Do you want to include result(s) for regexp submatch(es) (aka. capturing groups)?
		- Index : Do you want the "index(es)" of the result(s) or the "value(s)" of the result(s)?

		Basically, Find[All | First][string | []byte]Match[with Submatch][Index | Value]
		- So all methods that don't include "String" uses T<[]byte>, else T<string> as arg/res
	*/

	first_finders := func() {
		re := regexp.MustCompile(`\w+(\.[a-z]+)$`)

		/* ---- First Match Value Finders ---- */
		fmt.Println(string(re.Find([]byte("my_file.go"))))
		fmt.Println(re.FindString("my_file.go"))

		fmt.Println("--------")

		/* ---- First Match Index Finders ----- */
		fmt.Println(re.FindIndex([]byte("my_file.go")))
		fmt.Println(re.FindStringIndex("my_file.go"))

		fmt.Println("--------")

		/* ---- First Match+Submatch Value Finders ----- */
		subm := re.FindSubmatch([]byte("my_file.go"))
		subm_string := make([]string, len(subm))

		for i, bt := range subm {
			subm_string[i] = string(bt)
		}

		fmt.Println(subm_string)
		fmt.Println(re.FindStringSubmatch("my_file.go"))

		fmt.Println("--------")

		/* ---- First Match+Submatch Index Finders */
		fmt.Println(re.FindSubmatchIndex([]byte("my_file.go")))
		fmt.Println(re.FindStringSubmatchIndex("my_file.go"))
	}

	all_finders := func() {
		re := regexp.MustCompile(`\w+(\.[a-z]+)`)

		/* ---- All Match Value Finders ---- */
		all_bt := re.FindAll([]byte("my_file.go should be in a folder containing go.mod"), -1)
		all_string := make([]string, len(all_bt))

		for j, bt := range all_bt {
			all_string[j] = string(bt)
		}
		fmt.Println(all_string)
		fmt.Println(re.FindAllString("my_file.go should be in a folder containing go.mod", -1))

		/* ---- All Match Index Finders ----- */
		// re.FindAllIndex()
		// re.FindAllStringIndex()

		/* ---- All Match+Submatch Value Finders ----- */
		// re.FindAllSubmatch()
		// re.FindAllStringSubmatch()

		/* ---- All Match+Submatch Index Finders ----- */
		// re.FindAllSubmatchIndex()
		// re.FindAllStringSubmatchIndex()
	}

	first_finders()
	fmt.Println("--------")
	all_finders()
}

func replacers() {
	/* Method variants:
	1. If the method is a "String" variant, it uses T<string> arg/res, else T<[]byte> for argument and/or result

	2. If the method is a "Literal" variant, "$" signs inside the replacement value, normally interpreted as in Regexp.Expand, are substituted directly without using Regexp.Expand.

	3. If the method is a "Func" variant, the replacement parameter is rather a function value that yields a replacement value (string or []byte, depending on variant 1)
	*/

	// Basic
	re := regexp.MustCompile(`\.\$\{ext\}`)

	replaced_bts := re.ReplaceAll([]byte("file.${ext}"), []byte(".go"))
	fmt.Println(string(replaced_bts)) // file.go

	replaced_str := re.ReplaceAllString("main.${ext}", ".c")
	fmt.Println(replaced_str) // main.c

	re2 := regexp.MustCompile(`\.\$\{ext=(\.[a-z]+)\}`)

	replaced_str_exp := re2.ReplaceAllString("main.${ext=.cpp}", "${1}")
	fmt.Println(replaced_str_exp) // main.cpp

	// Other variants as described above
	replaced_bts_lit := re2.ReplaceAllLiteral([]byte("main.${ext=.cpp}"), []byte("${1}"))
	fmt.Println(replaced_bts_lit) // main${1}

	replaced_str_lit := re2.ReplaceAllLiteralString("main.${ext=.cpp}", "${1}")
	fmt.Println(replaced_str_lit) // main${1}

	// re.ReplaceAllFunc()
	// re.ReplaceAllStringFunc()
}

func regexp_others() {
	re := regexp.MustCompile(`(?P<fname>\w+)(?P<ext>\.[a-z]+)`)

	/*
		- SubexpNames(), returns the names of the "named" capturing groups in regexp
		- The name for the first sub-expression is names[1] and so on. names[0] is always an empty string so you don't need it
		- Note: The returned slice should not be modified

		- A variant, SubexpIndex(name string), returns the index of a specified name
		- Another variant, NumSubexp(), returns the number of capturing groups in the regexp, doesn't consider that first empty string
	*/

	fmt.Println(re.SubexpNames())      // ["", fname, ext]
	fmt.Println(re.SubexpIndex("ext")) // 2
	fmt.Println(re.NumSubexp())        // 2

	fmt.Println("--------------")

	fmt.Println(re.String()) // returns the compiled regex (re) string

	fmt.Println("--------------")

	re_non_digit := regexp.MustCompile(`\s*\D\s*`)
	// this is analogous to splitting with regex in javascript
	// this is also analogous to strings.SplitN provided the lookup is an exact string
	operands := re_non_digit.Split("1 + 2 * 3 / 4 - 5", -1 /* number of substrings to return */)
	fmt.Println(operands) // [1 2 3 4 5]

	// this is equivalent to using the opposite of `re_sp`'s pattern `\d` with FindAllString()

	re_digit := regexp.MustCompile(`\d`)
	operands2 := re_digit.FindAllString("1 + 2 * 3 / 4 - 5", -1)
	fmt.Println(operands2) // [1 2 3 4 5]

	fmt.Println("-----------------")

	lp_cf := regexp.MustCompile(`Good (morning|afternoon|evening)`)
	lp_ct := regexp.MustCompile(`Good morning`)

	// prefix: returns a literal string that must begin any match of the regexp
	// complete: does the literal string consist of the entire regexp?
	prefix_cf, complete_cf := lp_cf.LiteralPrefix()
	prefix_ct, complete_ct := lp_ct.LiteralPrefix()

	fmt.Println(prefix_cf, complete_cf) // "Good ", false
	fmt.Println(prefix_ct, complete_ct) // "Good morning", true

	/* others */
	// re.MarshalText() // implements encoding.TextMarshaler interface
	// re.UnmarshalText() // implements encoding.TextUnmarshaler interface
	// re.Longest()
	// re.FindReaderIndex()
	// re.FindReaderSubmatchIndex()

}
