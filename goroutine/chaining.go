package main

import (
	"flag"
	"fmt"
	"runtime"
)

var ngrountine = flag.Int("n", 1000, "how many goroutines")

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngrountine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0      // start the chaining
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 100000, approx. 1,5 s
}
