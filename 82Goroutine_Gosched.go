package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	t1 := time.Now()
	exit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("go:", i)
		}
		exit <- true
	}()

	go func() {
		for {
			fmt.Println("go2")
		}
	}()

	<-exit
	t2 := time.Now()

	fmt.Println("tack time:", t2.Sub(t1))

}
