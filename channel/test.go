package main

import (
	"fmt"
	"time"
)

var (
	tickets int = 1000
	person  int = 100000
)

func sigleChan() {
	start := time.Now()
	for i := 0; i < person; i++ {
		if tickets <= 0 {
			fmt.Printf("票已经卖完\n")
			break
		}
		tickets = tickets - 1
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("person %d 获取了一张票\n", i)
	}
	fmt.Printf("所有时间：%v\n", time.Now().Sub(start))
}

func mutilChan() {
	chans := make([]chan int, 10)
}

func main() {
	sigleChan()
}
