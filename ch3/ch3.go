package main

import (
	"fmt"
	// "unicode/utf8"
)

/* Runes, Unicode, UTF-8 */
/* func main() {
	s := "✨😎🎯🎈"
	rSlice := []rune(s)
	rChar := rune('✨')
	fmt.Printf("%x\n", s) // force-breaking to ASCII
	fmt.Printf("%x\n", rSlice)
	fmt.Printf("%U\n", rSlice)
	fmt.Println(string(rSlice))
	fmt.Printf("%U\n", rChar)
} */

/* func comma(numString string) string {
	var buf bytes.Buffer
	for i, d := range numString {
		fmt.Fprintf(&buf, "%c", d)
		if (i+1)%3 == (len(numString)%3) && (i+1) != len(numString) {
			buf.WriteString(", ")
		}
	}
	return buf.String()
} */

type Weekday int

const (
	Sunday Weekday = 1 + iota
	Monday
	Tuesday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

type DataUnit int

const (
	KB DataUnit = (1 << iota) * 1024
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const x int = 5

var arr = [x]int{1, 2, 3, 4, 5}

func main() {
	fmt.Printf("%T %[1]v\n", MB)
}
