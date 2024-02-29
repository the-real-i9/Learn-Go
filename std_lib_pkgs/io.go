package main

/*
- The io package provides basic interfaces to I/O primitives. Its primary job is to wrap existing implementations of such primitives, such as those in package os, into share public interfaces that abstract the functionality.
- It provides several strategies to perform read/write operations on I/O.
- It provides a common interface to manipulate any I/O through abstracion, the common one being io.Writer's Write() and io.Reader's Read()
	- Every I/O component implements one or more of its interfaces as it applies to it.
	- Therefore, we can manipulate any I/O using the interfaces provided and the underlying implementation of the concrete type will run.
*/

func io_pkg() {
	// reads data from io.Reader (input), and writes them into io.Writer (output)
	// io.Copy()

	//like io.Copy() except that it stages througn the provided buffer (if one is required) rather than allocating a temporary one.
	// io.CopyBuffer()

	// explicitly copy n bytes from io.Reader (input) to io.Writer (output)
	// io.CopyN()

	// io.Pipe() creates a "synchronous", "unbuffered", in-memory pipe,
	// it returns a read end (*io.PipeReader) and a write end (*io.PipeWriter) to manipulate the ends of the pipe as usual
	// both ends implements ReadCloser or WriteCloser as applicable, both closer have a CloseWithError(err) variant.
	// read_end, write_end := io.Pipe()

	// read all and return ([]byte) read data from io.Reader(), until io.EOF
	// io.ReadAll()

	// read at least min bytes into the provided buffer
	// io.ReadAtLeast()

	// read into the provided buffer, bytes equal to its length
	// io.ReadFull()

	// write the provided string into io.Writer (output)
	// io.WriteString()

	// modify an io.Reader to one with a maximum readable byte
	// any read on the io.Reader will stop after the maximum readable byte has been read
	/* lim_reader := io.LimitedReader{R: strings.NewReader("Hey there"), N: 5}
	lim_reader2 := io.LimitReader(strings.NewReader("Hey there"), 7)

	io.Copy(os.Stdout, &lim_reader) // Hey t
	io.Copy(os.Stdout, lim_reader2) // Hey the */

	// create a single io.Reader from multiple io.Reader(s) by concatenating their content
	// reading from the resulting reader, a single concatenated data
	// mult_reader := io.MultiReader()

	// write to multiple io.Writer(s) at the same time, from a single io.Writer returned
	// mult_writer := io.MultiWriter()

	// attach to an io.Reader, an io.Writer, such that
	// any reads on the modified io.Reader automatically writes to the io.Writer
	// tee_reader := io.TeeReader()

	// create a SectionReader from a Reader (implementing ReaderAt), that can
	// read from a specified section of the Reader's data as desired
	// the returned SectionReader is a Reader containing n bytes from the offset as specified
	// sec_reader := io.NewSectionReader()

	// create an OffsetWriter from a Writer (implenting WriterAt), that can
	// write to a specified section of the Writer as desired.
	// the returned OffsetWriter is a Writer that always writes data starting from the specified offset
	// off_writer := io.NewOffsetWriter() // notice that it is analogous to SectionReader for a Reader

}
