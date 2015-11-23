package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("sync:", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("Buffered:", testing.Benchmark(BenchmarkChannelBuffered).String())
}

func BenchmarkChannelSync(t *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < t.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for _ = range ch {

	}
}

func BenchmarkChannelBuffered(t *testing.B) {
	ch := make(chan int, 256)
	go func() {
		for i := 0; i < t.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for _ = range ch {

	}
}
