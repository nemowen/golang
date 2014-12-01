package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	println("start:", runtime.NumGoroutine()) // NumGoroutime : 等待排队的Goroutine数量

	for i := 0; i < 10; i++ {
		go func(n int) {
			println(n, runtime.NumGoroutine())
		}(i)
	}

	time.Sleep(3 * time.Second)
	println("over:", runtime.NumGoroutine())
}
