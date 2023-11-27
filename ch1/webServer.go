package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex // a mutex lock for synchronization
var count int     // a shared variable

func main() {
	http.HandleFunc("/", handler) // handle GET calls to "/"
	http.HandleFunc("/count", counter)

	// listener
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(res http.ResponseWriter, req *http.Request) {
	mu.Lock()   // acquire lock
	count++     // modify shared variable
	mu.Unlock() // release lock
	fmt.Fprintf(res, "<p>This is a page at %q. This is from Kenny</p>\n", req.URL.Path)
}

func counter(res http.ResponseWriter, req *http.Request) {
	mu.Lock()
	fmt.Fprintf(res, "<p>Count %d</p>", count)
	mu.Unlock()
}
