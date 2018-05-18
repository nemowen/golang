package main

import (
	"fmt"
	"time"
)

func timer(d time.Duration, do func()) <-chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(d)

		do()

		ch <- 1
	}()

	return ch
}

func do() {
	fmt.Println("done")
}

func main() {
	for i := 0; i <= 24; i++ {
		ch := timer(2000*time.Microsecond, do)

		fmt.Println(<-ch)
	}
}
