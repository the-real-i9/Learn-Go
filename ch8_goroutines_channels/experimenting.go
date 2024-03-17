package main

import (
	"fmt"
	"time"
)

var postBoxes = make(map[string]chan int)

func sender() {
	for i := 1; ; i++ {
		time.Sleep(1 * time.Second)
		postBoxes["box1"] <- i
	}
}

func trigger() {
	go sender()
}

func Experimenting() {
	go func() {
		postBoxes["box1"] = make(chan int)
	}()

	trigger()

	time.Sleep(3 * time.Second)

	for j := range postBoxes["box1"] {
		fmt.Println(j)
	}

}
