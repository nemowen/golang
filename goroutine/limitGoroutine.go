package main

import (
	"fmt"

	"sync"
	"time"
)

const LIMIT_GO_NUM = 2

var wg sync.WaitGroup
var ids = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func worker(id int, limit chan struct{}) {
	fmt.Println(id)
	//println("Goroutine num:", runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	wg.Done()
	<-limit
}

func main() {
	wg.Add(len(ids))
	limit := make(chan struct{}, LIMIT_GO_NUM)
	for _, id := range ids {
		limit <- struct{}{}
		go worker(id, limit)
	}
	wg.Wait()
}
