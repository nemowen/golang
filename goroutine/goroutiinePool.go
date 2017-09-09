package main

import (
	"fmt"
	"runtime"
	"strconv"
)

const (
	WORKER_COUNT = 10
	ID_COUNT     = 100
)

type Message struct {
	id     int
	result string
}

// 不需要频繁创建和销毁goroutine，创建一个goroutine池更合适。
func main() {
	chId := make(chan int, WORKER_COUNT)
	chResult := make(chan Message, WORKER_COUNT)
	for i := 0; i < WORKER_COUNT; i++ {
		go func() {
			for {
				id := <-chId
				chResult <- Message{
					id:     id,
					result: strconv.Itoa(id),
				}
			}
		}()
	}
	go func() {
		for i := 0; i < ID_COUNT; i++ {
			chId <- i
		}
	}()
	for i := 0; i < ID_COUNT; i++ {
		msg := <-chResult
		fmt.Printf("id[%d]=>result[%s] : %d\n", msg.id, msg.result, runtime.NumGoroutine())
	}
}


