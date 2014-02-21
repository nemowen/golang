package main

import (
	"fmt"
	"runtime"
)

func test(c chan bool, index int) {
	var a, i int64
	for i = 0; i < 10000000; i++ {
		a += i
	}
	//fmt.Println(index, ":", a)
	if 99 == index {
		c <- true
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan bool, 1)
	for i := 0; i < 100; i++ {
		test(c, i)
	}

	<-c

	fmt.Println("all done...")

}
