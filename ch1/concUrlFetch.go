package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine

		/*
			- For each url, the `go` statement in this loop starts a new goroutine that calls `fetch` asynchronoushly.
				- Meaning, it doesn't block this loop (main routine), and control is immediately transferred to the next statement after this loop finishes.
			- All the goroutines started by this loop execute concurrently.
		*/
	}

	for range os.Args[1:] {
		/* (continued from `func fetch`) ... another goroutine (main) attempts the corresponding receive on channel `ch` */

		fmt.Println("I'm in the next loop")
		fmt.Println(<-ch)

		/*
			- This loop (main routine) blocks (when it attempts to receive and) until it receives from channel `ch`.
			- It blocks/waits for data to arrive on the channel `ch`, when it gets one, it immediately consumes it (zero capacity buffer | unbuffered channel | no temporary storing)
			- It repeats the loop (up to the number of urls), and wait for another data (that is expected) to arrive on channel `ch`
		*/
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
	/* When each goroutine (fetch) attempts a send on channel `ch`, it blocks, until ... (continued in `func main`)*/

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
