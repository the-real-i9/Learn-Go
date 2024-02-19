package main

import (
	"fmt"
	"math"
)

/*
+ When a channel is supplied as a function parameter, it is nearly always with the intent that it be used exclusively for "sending" or for "receiving".

+ To document this intent and prevent misuse, the Go type system provides "unidirectional" channel types that expose only one or the other of the send and receive operations.
	- The type `chan<- int`, a "send-only" channel of `int`, allows sends but not receives.
	- The type `<-chan int`, a "receive-only" channel of `int`, allows receives but not sends.
*/

func counter(out chan<- int) {
	for i := 0; i <= 32; i++ {
		out <- i
	}
	close(out)
}

func bitsequencer(out chan<- int, in <-chan int) {
	for j := range in {
		out <- int(math.Pow(2, float64(j)))
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func pipel_4() {
	naturals := make(chan int)
	bitseq := make(chan int)

	// Conversion from bidirectional to unidirectional channel types are permitted in any assignment

	go counter(naturals)
	go bitsequencer(bitseq, naturals)
	printer(bitseq)
}

func unidirectionalChan() {
	pipel_4()
}
