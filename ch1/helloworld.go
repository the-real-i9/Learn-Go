package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	// "learn/helpers" // an imported custom package (module)
)

func main() {
	/* url := "https://www.facebook.com"
	resp, err := http.Get(url) */
	// b, err := io.ReadAll(resp.Body)
	// w, err2 := io.Copy(os.Stdout, resp.Body)
	/* file, ferr := os.Create("resp_body.html")
	_, ferr2 := io.Copy(file, resp.Body)
	resp.Body.Close()
	file.Close() */
	/* -------------------------- */
	/* -------------------------- */

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
		/* For each url, the `go` statement in the first range loop starts a new goroutine that calls `fetch` asynchronoushly */
		/* When a goroutine (fetch) attempts a send on our (ch) channel, it blocks, until ... */
	}
	for range os.Args[1:] {
		/* ... another goutine (main) attempts the corresponding receive on our channel */
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
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
