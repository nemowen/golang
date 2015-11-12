package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= 9; i++ {
		jobs <- i
	}

	//close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
}
