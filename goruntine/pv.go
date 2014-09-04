package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	exit := make(chan bool)
	depot := make(chan int, 5) // 仓库:最多能存5个产品

	// 一个进程生产产品
	go func() {
		rand.Seed(time.Now().Unix())
		for i := 1; i <= 100; i++ {

			// 生产后放入到仓库
			depot <- rand.Intn(1000000)

		}
		close(depot)
	}()

	// 一个进程消费产品
	go func() {
		for a := range depot {
			fmt.Print(a, " ")
		}
		exit <- true
	}()

	<-exit
}
