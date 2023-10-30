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

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()   // acquire lock
	count++     // modify shared variable
	mu.Unlock() // release lock
	fmt.Fprintf(w, "<p>This is a page at %q</p>", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "<p>Count %d</p>", count)
	mu.Unlock()
}
