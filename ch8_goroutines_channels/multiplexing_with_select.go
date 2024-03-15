package main

import (
	"fmt"
	"os"
	"time"
)

/*
+ Idea: Select one of several channels, of which data arrives at "first"

+ The select statement:
	- Like a switch statement, it has a number of cases and an optional `default`.
	- Each case specifies a communication(a send or receive operation on some channel) and an associated block of statements.

+ A `select` waits until a communication for some case is ready to proceed. It then performs that communication and executes the case's associated statements; the other commnunications do not happen. If multiple cases are ready, `select` picks one at random, which ensures that every channel has an equal chance of being selected.

+ Sometimes we want to try to send or receive on a channel but avoid blocking if the channel is not ready, a "non-blocking" communication.
			- A `select` may have a `default`, which specifies what to do when none the other communications can proceed immediately.

+ The zero value for a channel is `nil`.
+ Because send and receive operations on a `nil` channel block forever, a case in a select statement whose channel is nil is never selected.
	- This lets us use `nil` to enable or disable cases that correspond to features like hasdling timeouts or cancellation, responding to other input events, or emitting output.
*/

/*
+ The time.Tick function returns a channel on which it sends events periodically
+ We add the ability to abort the launch countdown by pressing the return key during the countdown.
+ The select statement below causes each iteration of the loop to wait up to 1 second for an abort, but no longer.
*/
func multiplexSelect() {
	abort := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- 1
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick.C:
			// proceed
		case <-abort:
			tick.Stop()
			fmt.Println("Launch aborted!")
			return
			// default: // polling both channels | a non-blocking receive
			// proceed -- this would cause the loop to run like normal, so fast that it completes before an event on the abort channel
		}
	}
}
