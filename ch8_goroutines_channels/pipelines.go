package main

import (
	"fmt"
	"math"
)

/*
+ Channels can be used to connect goroutines together so that the output of one is the unput to another. This is called a "pipeline".
*/

func pipelines() {
	pipel_3()
}

/* The first goroutine, "counter", generates the integers 0, 1, 2, ..., and sends them over a channel to the second goroutine, "squarer", which receives each value, multiply it by 2, and sends the result over another channel to the third goroutine, "printer", which receives the squared values and prints them. */
func pipel_1() {
	naturals := make(chan int)
	bitseq := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x // send x to "naturals channel"
		}
	}()

	// squarer
	go func() {
		for {
			x := <-naturals
			bitseq <- int(math.Pow(2, float64(x)))
		}
	}()

	// main: printer
	for {
		fmt.Println(<-bitseq)
	}
}

/*
+ pipel_1() runs forever. But what if we want to send only a finite number of values through the pipline???

+ If the sender knows that no further values will ever be sent on a channel, it is useful to communicate this fact to the receiver goroutines so that they can stop waiting.
	- This is accomplished by closing the channel using the built-in `close()` function.

+ After a channel has been closed, any further send operations on it will panic. After the closed channel has been drained, that is, after the last sent element has been received, all subsequent receive operations will proceed without blocking but will yield a zero value.
*/

func pipel_2a() {
	naturals := make(chan int)
	bitseq := make(chan int)

	// counter
	go func() {
		for x := 0; x <= 32; x++ {
			naturals <- x // send x to "naturas channel"
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for {
			// this will eventually be an operation on a closed channel that forever yields zeros
			x := <-naturals
			bitseq <- int(math.Pow(2, float64(x)))
		}
	}()

	// main: printer
	for {
		fmt.Println(<-bitseq)
	}
}

/*
+ There is no way for a reveiver goroutine to test directly whether a channel has been closed, but there is a variant of the receive operation that produces two results: the received channel element, plus a boolean value, conventionally called `ok`, which is `true` for a successful receive and `false` for a receive on a closed and drained channel.
*/
func pipel_2b() {
	naturals := make(chan int)
	bitseq := make(chan int)

	// counter
	go func() {
		for x := 0; x <= 32; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok { // after the `naturals` channel is drained
				break
			}
			bitseq <- int(math.Pow(2, float64(x)))
		}
		close(bitseq)
	}()

	// main: printer
	for {
		x, ok := <-bitseq
		if !ok { // after the `bitseq` channel is drained
			break
		}
		fmt.Println(x)
	}
}

/*
+ Because the syntax above is clumsy and this pattern is common, the language lets us use a `range` loop to iterate over channels too. This is a more convenient syntax for receiving all the values sent on a channel and terminating the loop after the last one.

+ You don't have to close every channel after you're finished with it. It's only necessary to close a channel when it is important to tell the receiving goroutines that all data have been sent.
*/

func pipel_3() {
	naturals := make(chan int)
	bitseq := make(chan int)

	// counter
	go func() {
		for x := 0; x <= 32; x++ {
			naturals <- x // send x to "naturas channel"
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for x := range naturals {
			bitseq <- int(math.Pow(2, float64(x)))
		}
		close(bitseq)
	}()

	// main: printer
	for x := range bitseq {
		fmt.Println(x)
	}
}
