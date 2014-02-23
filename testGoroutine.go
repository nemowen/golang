package main

import (
	"fmt"
	"runtime"
	"time"
)

func test(c chan bool, b bool) {
	x := 0
	for i := 0; i < 100000000; i++ {
		x += i
	}
	fmt.Println(x)
	if b {
		close(c)
	}
}

func main() {
	stime := time.Now()
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	c := make(chan bool)
	for i := 0; i < 100; i++ {
		test(c, i == 99)
	}

	<-c

	fmt.Println("all done...")
	etime := time.Now()

	fmt.Println("tack time:", etime.Sub(stime))

}
