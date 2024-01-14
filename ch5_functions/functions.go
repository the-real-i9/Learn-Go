package main

import (
	"fmt"
	"io"
	"net/http"
)

func getWebPage(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%v %v", url, err)
	}

	// defer resp.Body.Close()
	return resp.Body, nil
}

func main() {
	/* webPage, err := getWebPage(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "getWebPage: %v\n", err)
		os.Exit(1)
	} */

	// findlinks1(webPage)
	// findtags(webPage)
	// countElements(webPage)
	funcVals()
}
