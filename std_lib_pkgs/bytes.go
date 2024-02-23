package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

/*
"bytes" functions (on byte slices) are analogous to "strings" functions
"bytes" `type Reader` are analogoues to "strings" `type Reader` also

You should be able to handle the `type Reader` as you've
already dealt with analogous reader functions from `type Buffer`
*/

func bytes_pkg() {
	bytes_pkg3()
}

/* ==== AvailableBuffer(), Reader family, Len(), Bytes(), String() */
func bytes_pkg1() {
	buf := bytes.Buffer{}
	// var buf bytes.Buffer
	// fmt.Println(buf.Available())
	for i := 0; i <= 5; i++ {
		// creates an empty (semantic, not actual) buffer ([]byte - empty byte slice)
		// with buf's capacity, like a temporary working buffer
		b := buf.AvailableBuffer()
		b = append(b, fmt.Sprintf("%d_", i)...)
		buf.Write(b)
	}

	// Bytes() returns a slice holding the "unread portion" of the buffer
	bts := buf.Bytes()
	fmt.Printf("Unread portion: %s\n", string(bts))

	// Len() returns the number of bytes of the "unread portion" of the buffer
	fmt.Printf("Unread portion's len: %d\n\n", buf.Len())

	// What's meant by "unread portion"? Functions that Read bytes from buffer consumes the buffer, as if the byte(s) read have been taken out of the buffer. If functions that do return the unread portion are executed afterwards, they'll only return the uread bytes. Let's see this for real:

	// Next() is a reader
	fmt.Printf("`Next(2)` reads 2 bytes: %s\n\n", string(buf.Next(2))) // 0_

	// String() also returns the unread portion
	fmt.Printf("Unread portion: %s\n", buf.String())      // 1_2_3_4_5_
	fmt.Printf("Unread portion's len: %d\n\n", buf.Len()) // 10
	// Next(n) has read n bytes away

	test_bf := make([]byte, 2)

	if _, err := buf.Read(test_bf); err != nil { // Read() is another reader
		panic(err)
	}
	fmt.Printf("`Read(make([]byte, 2))` reads 2 bytes: %s\n\n", string(test_bf)) // 1_
	fmt.Printf("Unread portion: %s\n", buf.String())                             // 2_3_4_5_
	fmt.Printf("Unread portion's len: %d\n\n", buf.Len())                        // 8
	// Read(p []byte)  has read another len(p) bytes away into p

	// ReadByte() reads the next byte
	// ReadRune() reads the next rune

	// ReadBytes(delim byte) reads until the first occurence of delim, returning the byte slice containing the data up to delim
	line_byt, err := buf.ReadBytes('_')
	if err != nil {
		panic(err)
	}
	fmt.Printf("ReadBytes('_') reads up to and including the first underscore: %s\n\n", string(line_byt)) // 2_
	fmt.Printf("Unread portion: %s\n", buf.String())                                                      // 3_4_5_
	fmt.Printf("Unread portion's len: %d\n\n", buf.Len())                                                 // 6

	// ReadString(delim byte) reads until the first occurence of delim, returning the string containing the data up to delim
	line_str, err := buf.ReadString('_')
	if err != nil {
		panic(err)
	}
	fmt.Printf("ReadString('_') reads up to and including the first underscore: %s\n\n", string(line_str)) // 3_

	fmt.Printf("Unread portion: %s\n", buf.String())    // 4_5_
	fmt.Printf("Unread portion's len: %d\n", buf.Len()) // 4
}

/* ==== Grow(), Cap() ==== */
func bytes_pkg2() {
	buf := bytes.Buffer{}
	fmt.Println(buf.Cap())

	buf.Grow(10) // must be called for the `bb` to have a finite capacity
	fmt.Println(buf.Cap())

	bb := buf.Bytes() // without buf.Grow(), `bb` here has capacity 0. It'd have been fine if this line had come after buf.Write() which will implicity grow the buffer.
	buf.Write([]byte("kenny"))
	// bb := buf.Bytes() // without buf.Grow() `bb` here is fine, as the previous line has implicitly grown `buf`'s capacity

	fmt.Println(string(bb[:buf.Len()])) // this operation will panic if `buf`'s capacity hasn't found a way to grow before accessing `bb`, since `bb`'s capacity is 0, it'd be out of bounds.
	// Notice that it works after calling buf.Grow() which grows the buf's capacity explicitly or buf.Write() which does the ssame thing implicitly.
	// Notice also that it successfully prints "kenny" even though `bb` came before buf.Write. This is because although `bb` has len 0, the slicing operation's bounds is actually accessing `bb`'s underlying array which is the same as that of `buf` that has been modified by buf.Write() afterwards.
}

/* ==== Write() family, ReadFrom(), WriteTo() ==== */
func bytes_pkg3() {
	buf := bytes.Buffer{}

	// ReadFrom(io.Reader) reads from io.Reader and appends it to buf
	buf.ReadFrom(strings.NewReader("Reading from *strings.Reader"))

	// the following appends to buf
	buf.Write([]byte("\nI'll be back"))
	buf.WriteByte(',')
	buf.WriteRune('💨')
	buf.WriteString(" in a flash\n")

	// WriteTo(io.Writer) empties the buffer contents into io.Writer
	buf.WriteTo(os.Stdout)
	fmt.Println(buf.String()) // "" // buf has been emptied
}

/* ===== UnreadByte(), UnreadRune(), Reset(), Truncate(), ReadRune() ===== */
func bytes_pkg4() {
	buf := bytes.Buffer{}
	buf.WriteString("Kenny")

	fmt.Println(buf.String()) // "Kenny"

	// reading the next 1 byte away. This could have been done by any buf.{reader} function
	fmt.Println(string(buf.Next(1))) // "K"
	fmt.Println(buf.String())        // "enny"

	// UnreadByte() unreads the last byte returned by the most recent successful read operation tht read at least one byte
	// unreading (restoring) the last successfully read byte (by Next(1), above)
	if err := buf.UnreadByte(); err == nil {
		fmt.Println(buf.String()) // "Kenny" // "K" is restored
	}
	// Note: UnreadByte() would only restore the single last read byte, even if `n` bytes were read

	buf.Reset() // resets (empties) the buffer (len(buf) = 0) but retains the buffer's capacity, same as buf.Truncate(0)
	buf.WriteString("🤑Samuel")

	rn, _, _ := buf.ReadRune()

	fmt.Println(string(rn))   // "🤑"
	fmt.Println(buf.String()) // "Samuel"

	// UnreadRune() is the rune equivalent of UnreadByte()
	if err := buf.UnreadRune(); err == nil {
		fmt.Println(buf.String()) // "🤑Samuel" // "🤑" is restored
	}
	// Note: UnreadRune() also would only restore the single last read rune, even if `n` runes were read

	// Truncate(n int) discards all except the first n unread bytes from the buffer, while retaining the buffer's capacity
	buf.Truncate(4 + 3)       // Note: the preceeding emoji is a 4byte unicode char, plus 3 characters (1 ascii char == 1 byte)
	fmt.Println(buf.String()) // "🤑Sam"
}
