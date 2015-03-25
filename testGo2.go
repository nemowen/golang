package main

import "fmt"
import "runtime"
import "time"

func run(i, n int, ch chan int) {
	count := 0
	for i := i; i < n; i++ {
		count = count + i
	}
	ch <- count
}

func main() {
	t1 := time.Now()
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)
	chs := make([]chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		chs[i] = make(chan int)
		n := 900000000 / NCPU
		go run(i*n, (i+1)*n, chs[i])
	}

	count := 0
	for i := 0; i < NCPU; i++ {
		t := <-chs[i]
		count = count + t
	}
	t2 := time.Now()

	fmt.Printf("cpu num:%d,cost:%s,count:%d\n", NCPU, t2.Sub(t1), count)
}
