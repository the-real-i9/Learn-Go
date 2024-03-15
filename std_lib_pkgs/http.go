package main

import (
	"fmt"
	"log"
	"net/http"
)

func http_pkg() {
	h1 := func(res http.ResponseWriter, _ *http.Request) {
		fmt.Println("Hi")
		res.Header().Set("hello", "How are you?")
	}

	h2 := func(res http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(res, "Good evening!"+res.Header().Get("hello"))
	}

	/* h3 := func(res http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(res, "How're you doing today?")
	} */

	// http.HandleFunc("/", h1)
	http.HandleFunc("/greet", h1)
	http.HandleFunc("/greet/question", h2)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
