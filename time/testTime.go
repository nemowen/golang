package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var num int = 1000
var lock = sync.Mutex{}
var exit = make(chan bool)
var ia int = 0
var ch = make(chan int, 1000)

func buyTicket(id int) {
	for {
		//lock.Lock()
		if num <= 0 {
			exit <- true
			break
		} else {
			ch <- 1
			fmt.Print(" [", id, ":", num, "]")
		}

	}
}

func main() {
	t0 := time.Now()
	runtime.GOMAXPROCS(1)

	for i := 0; i < 100; i++ {
		go buyTicket(i)
	}
	for i := 0; i < 1000; i++ {
		<-ch
		num--
		ia++
	}
	<-exit
	t1 := time.Now()
	fmt.Println("\n----------------", ia)
	fmt.Println("time:", t1.Sub(t0))
}
