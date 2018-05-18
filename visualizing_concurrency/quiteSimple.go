package main

import "fmt"

func main() {
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 1 to channel
		ch <- 1
	}()

	// read from channel and print
	fmt.Println(<-ch)
}
