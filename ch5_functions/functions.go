package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func findlinks1(webPage io.Reader) {
	doc, err := html.Parse(webPage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func findtags(webPage io.Reader) {
	doc, err := html.Parse(webPage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findtags1: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc)
}

func getWebPage(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%v %v", url, err)
	}

	// defer resp.Body.Close()
	return resp.Body, nil
}

func main() {
	webPage, err := getWebPage("https://golang.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "getWebPage: %v\n", err)
		os.Exit(1)
	}

	findlinks1(webPage)
	// findtags(webPage)
}
