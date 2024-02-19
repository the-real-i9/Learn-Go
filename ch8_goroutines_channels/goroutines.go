package main

import (
	"fmt"
	"time"
)

func goRoutines() {
	concurrentProgram()
	sequentialProgram()
}

/*
When a program starts, its only goroutine is the one that calles the main function, called the "main goroutine"
New goroutines are created by the `go` statement: An ordinary function or method call prefixed with the `go` keyword.

A `go` statement causes the function to be called in a newly created goroutine.
The go statement itself completes immediately

Newly created goroutines makes progress concurrently with every other running goroutine.
*/

func concurrentProgram() {
	// But in most cases we don't want to wait, as we don't need the result
	// or we might want to do other things while we wait for the result
	fmt.Println("------------------")
	fmt.Println("Concurrent program")
	fmt.Println("------------------")
	fmt.Println("CP: Process 1 will run")
	go slowProcess1("CP: ")
	fmt.Println("CP: Process 2 will run")
	go slowProcess2("CP: ")
}

func sequentialProgram() {
	// In a sequential program the next procedure has to wait for the previous one to finish
	fmt.Println("------------------")
	fmt.Println("Sequential program")
	fmt.Println("------------------")
	slowProcess1("SP: ")
	slowProcess2("SP: ")
}

func slowProcess1(p string) {
	fmt.Println(p + "Process 1 has started")
	time.Sleep(5 * time.Second)
	fmt.Println(p + "Process 1 has finished")
}

func slowProcess2(p string) {
	fmt.Println(p + "Process 2 has started")
	time.Sleep(5 * time.Second)
	fmt.Println(p + "Process 2 has finished")
}
