package main

import (
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	tick := time.Tick(1 * time.Second)     // 秒针滴答声
	timeout := time.After(5 * time.Second) // 设置超时机制

	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-timeout:
				println("timeout")
				//o <- true
				break
			case t := <-tick:
				println("tick:", t.Second())
			}

		}
	}()
	c <- 1
	c <- 3
	<-o
}
