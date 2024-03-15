package main

import (
	"log"
	"time"
)

/*
+ A send operation on an ubuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel. If the receive operation was attempted first, the receiving goroutine blocks until another goroutine performs a send on the same channel.

+ Communication over an ubuffered channel causes the sending and receiving goroutines to synchronize. Consequentely, they are sometimes called synchronous channels.

+ Sometimes we just want to use channels for synchronization purposes without needing their value. In such cases any simple and conventional value to represent content is used
*/

func unbuffChannels() {
	unbuff_chan := make(chan int)

	// sending goroutine
	go func() {
		log.Print("Asleep...")
		time.Sleep(1 * time.Second)
		log.Print("Awake")
		unbuff_chan <- 1 // signal the main goroutine
	}()

	// main: receiving goroutine
	// this will definitely be attempted first (as it is the main goroutine) and would, consequently, block until the sending goroutine above sends to it

	// fmt.Println(<-unbuff_chan) // if value is needed

	// "wait for background goroutine to finish"
	<-unbuff_chan // for mere syncronization
}
