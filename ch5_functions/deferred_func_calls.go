package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func add(x int, y int) {
	fmt.Println(x + y)
}

func sub(x int, y int) {
	fmt.Println(x - y)
}

func div(x int, y int) {
	fmt.Println(x / y)
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index"
	}
	local = local + ".html"

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return local, n, err
}

func defFuncCalls() {
	defer add(10, 5)
	defer sub(10, 5)
	add(20, 14)
	sub(20, 14)
}
