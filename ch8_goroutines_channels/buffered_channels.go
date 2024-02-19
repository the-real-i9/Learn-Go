package main

import (
	"fmt"
	"time"
)

/*
+ A buffered channel has a queue of elements. The queue's maximum size is determined when it is created, by the capacity argument to `make`.

+ A send operation on a buffered channel inserts an element at the back of the queue, and a receive operation removes an element from the front.

+ If the channel is full, the send operation blocks its goroutine until space is made available by another goroutine's receive. Conversely, if the channel is empty, a receive operation blocks until a value is sent by another goroutine.

+ The `cap(ch)` operation returns the capacity of a channel's buffer. The `len(ch)` operation returns the number of elements currently buffered.

+ Don't use buffered channels within a single goroutine as a queue, this is a mistake. If all you need is a simple queue, make one using a slice.

+ Unbuffered channels give stronger synchronization guarantees because every send operation is synchronized with its corresponding receive; with buffered channels, these operations are decoupled.
	- When we know an upper bound on the number of values that will be sent on a channel, it's not unusual to create a buffered channel of that size and perform all the sends before the first value is received.
	- Failure to allocate sufficient buffer capacity would cause the program to deadlock.sc
*/

func buffChan() {
	ch := make(chan string, 3)

	// we can send up to three values on this channel without the goroutine blocking
	ch <- "A"
	ch <- "B"
	ch <- "C"

	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println(fastestCake())
}

/*
+ The function below, makes parallel requests from 3 cake makers. It sends their readyCake over a buffered channel, then receives and returns only the first readyCake, which is the quicket to arrive.

+ Had we used an unbuffered channel, the two slower goroutines would have gotten stuck trying to send their readyCakes on a channel from which no goroutine wil ever receive. This situation is called a "goroutine leak", a bug.
*/

func fastestCake() string {
	readyCake := make(chan string, 3)
	go func() { readyCake <- makeCake("Maker 1", 1*time.Second) }()
	go func() { readyCake <- makeCake("Maker 2", 2*time.Second) }()
	go func() { readyCake <- makeCake("Maker 3", 3*time.Second) }()
	return <-readyCake
}

func makeCake(maker string, d time.Duration) string {
	time.Sleep(d)
	return maker + " done!"
}
