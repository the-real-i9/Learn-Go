package main

func strings_pkg() {
	/* === Clone() === */
	/* str := "Kenny"
	// strcl_np := strings.Clone(str) // not preferred

	strcl_p := str[:] // preferred

	fmt.Println(strcl_p) */

	/* ====== Compare() ====== */
	/* str1 := "Kenny"
	str2 := "Samuel"
	// cmp_int := strings.Compare(str1, str2)

	cmp_res := str1 < str2 // preferred: >, <, ==

	fmt.Println(cmp_res) */

	/* ====== Contains() family ======= */
	/* str := "Banana"
	fmt.Println(strings.Contains(str, "na"))                  // true, contains "na"
	fmt.Println(strings.Contains(str, "nx"))                  // false, does not contains "nx"
	fmt.Println(strings.ContainsAny(str, "nx"))               // true, contains "n"
	fmt.Println(strings.ContainsRune(str, 'B'))               // true, contains rune 'B'
	fmt.Println(strings.ContainsFunc(str, func(r rune) bool { // true, contains 'b' if 'B' is lowercased
		return unicode.ToLower(r) == 'b'
	})) */

	/* ===== Reader ===== */
	/* // the *strings.Reader type returned implements the io.Reader interface
	strReader := strings.NewReader("this is the content of a string reader")

	// io.Copy(os.Stdout, strReader) // the *strings.Reader type returned implements the io.Reader interface

	scanner := bufio.NewScanner(strReader)

	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	} */

	/* ===== more ==== */
	/* test_str := "Banana"

	fmt.Println(strings.Count(test_str, "na")) // 2

	before, after, found := strings.Cut(test_str, "a") // cuts the string at the first instance of "a"
	// before == "B"
	// after == "nana"
	// found == true
	fmt.Println(before, after, found)
	strings.CutPrefix(test_str, "B") // after == "anana"
	strings.CutSuffix(test_str, "na") // before == "Bana"
	strings.Map()
	strings.Join()
	strings.Index()
	strings.HasPrefix()
	strings.EqualFold()
	strings.Fields()
	strings.FieldsFunc() */

	/* ==== Types === */
	// strings.Builder -- for building strings
	// The methods of `Builder` are familiar

	// strings.Replacer -- for mutating strings
	/* rplc := strings.NewReplacer("s", "n")
	fmt.Println(rplc.Replace("basasa"))     // "banana"
	rplc.WriteString(os.Stdout, "basasa\n") // prints "banana" */

	// strings.Reader -- a string Reader
	// The methods of `Reader` are familiar

}
