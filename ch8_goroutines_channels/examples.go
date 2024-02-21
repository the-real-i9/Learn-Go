package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

func examples() {
	// concClockServer()
	// concEchoServer()
	concSum()
}

func concSum() {
	numCPUs := runtime.NumCPU()

	sum := sumSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 43, 34, 63, 2, 5, 2, 2, 5, 2, 5, 6, 3, 5, 2, 6, 2, 3, 6, 6, 2, 6, 2, 6, 2, 1, 4, 6, 8, 8}, numCPUs)

	fmt.Printf("Sum: %d -- Threads/CPUs: %d\n", sum, numCPUs)
}

func sumSlice(intSlice []int, threads int) int {
	chunkSum := make(chan int, threads)

	for t, x := threads, 0; t > 0; t-- {

		size := int(len(intSlice[x:]) / t)

		var chunk []int
		if len(intSlice[x:]) > size {
			chunk = intSlice[x : x+size]
		} else {
			chunk = intSlice[x:]
		}

		x += size

		go func() {
			var subSum int
			for _, y := range chunk {
				subSum += y
			}
			chunkSum <- subSum
		}()

	}

	var totalSum int
	for i := 0; i < threads; i++ {
		totalSum += <-chunkSum
	}

	return totalSum
}

func chunkSlice(intSlice []int, equal int) [][]int {
	var chunkRes [][]int
	size := int(len(intSlice) / equal)
	for x := 0; ; x += size {
		if len(intSlice[x:]) > size {
			chunkRes = append(chunkRes, intSlice[x:x+size])
		} else {
			chunkRes = append(chunkRes, intSlice[x:])
			break
		}
	}

	return chunkRes
}

/* --------------- */

func concEchoServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleEchoConn(conn)
	}
}

func handleEchoConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 3*time.Second)
	}
	conn.Close()
}

func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

/* --------------- */

func concClockServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleClockConn(conn)
	}
}

func handleClockConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("3:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
