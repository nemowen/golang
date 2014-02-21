package main

import (
	"fmt"
)

func fibonacci(c chan int64) {
	var x, y int64 = 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		}

	}
}
func main() {
	c := make(chan int64)

	//quit := make(chan int)
	go fibonacci(c)

	for i := 0; i < 10000; i++ {
		fmt.Printf("i:%d --- value:%d\n", i, <-c)
	}
	//quit <- 0

	fmt.Println("done...")

}
