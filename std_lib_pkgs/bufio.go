package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

/*
- It contains the Reader, Writer and Scanner interface types with their usual methods (those already familiar)
- The Reader and Writer interfaces consists familiar functions from the bytes.Buffer interface.
*/
func bufio_pkg() {
	/* Reader */
	// bufio.NewReader()

	/* Writer */
	// bufio.NewWriter()

	/* ReadWriter */
	// whatever you read from bufio.Reader automatically writes to bufio.Writer
	// like io.TeeReader
	read_writer := bufio.NewReadWriter(
		bufio.NewReader(strings.NewReader("How are you?")),
		bufio.NewWriter(os.Stdout),
	)

	read_writer.WriteString("Hello\n")

	rw_bt, _ := io.ReadAll(read_writer)

	os.Stdout.Write(rw_bt)

	/* Scanner: You already know how scanner works */
	// scanner := bufio.NewScanner(r io.Reader)

	// scanner.Split(split SplitFunc) -- What should be read per scan? Byte, Line (default), Rune, or Word.
	// scanner.Scan() -- Each call scans the next data as determined by SplitFunc
	// scanner.Bytes() -- data returned is byte based
	// scanner.Text() -- data returned is text based
	// scanner.Buffer() -- use an explicit buffer, rather than the default internal buffer

	// split functions
	// bufio.ScanBytes
	// bufio.ScanLines
	// bufio.ScanRunes
	// bufio.ScanWords
}
