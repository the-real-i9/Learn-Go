package main

import "fmt"

/* Channels: A message passing IPC mechanism
+ A channel is a communication mechanism that lets one goroutine send values to another goroutine.

+ Each channel is a conduit for values of a particular type, called the channel's element type.

+ A channel is a reference type... and you know what that means...

+ A channel has two principal operations, send and receive, collectively known as communications.
	- a send statement transmits value (of the channel's type) from one goroutine, through the channel, to another goroutine executing a corresponding receive expression

+ Channels support a third operation, close, after which no more values must be sent on this channel; subsequent attempts to send will cause a panic.
	- Receive operations on a closed channel yield the values that have been sent until no more values are left; any receive operations thereafter completes immediately and yields the zero value of the channel's element type.

+ Major types of channels
	unbuff_chan = make(chan int)        // unbuffered channel
	unbuff_chan_alt = make(chan int, 0) // unbuffered channel is basically a zero-capacity buffer channel
	buff_chan = make(chan int, 3)       // buffered channel with a buffer capacity of 3
*/

func channels() {
	// to create a channel
	ch := make(chan int) // unbuffered channel

	go func() {
		// ch <- {value} // a send statement
		// the value can be gotten anyhow, here we used a literal, it could've from anything that yields a value
		ch <- 5

		// to close a channel
		close(ch)
	}()

	// <-ch // a receive statment
	// the received value can be handled anyhow, here we've stored it in `r`, we could've used it directly anywhere a value is usable
	r := <-ch
	fmt.Println(r)
}
