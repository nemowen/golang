package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func main() {
	start := time.Now()
	DoMath(6, 4)
	fmt.Println("time:", time.Now().Sub(start))
}

func DoMath(Max, Num int) {

	filesMap := make([][]int, Num)
	chanList := []<-chan int{}

	for i := 0; i < Num; i++ {
		chanList = append(chanList, ChanInt(i+1))
		filesMap[i] = []int{}
	}

	wait.Add(Num)

	for j := 0; j < Num; j++ {

		go func(j int) {
			for i := 0; i < Max; i++ {
				filesMap[j] = append(filesMap[j], <-chanList[(j+i)%Num])
			}
			wait.Done()
		}(j)

	}

	wait.Wait()

	for k, v := range filesMap {
		fmt.Printf("%s : %v\n", string(rune(65+k)), v)
	}

}

func ChanInt(i int) <-chan int {

	ch := make(chan int, 0)

	go func() {
		for {
			ch <- i
		}
	}()

	return ch
}
