package main

import (
	"fmt"
	"sync"
)

const STEPS, NUMS = 6, 4

var wg sync.WaitGroup

func main() {
	files := make([]*File, 4)
	chans := []<-chan int{}

	for i := 0; i < NUMS; i++ {
		chans = append(chans, newChan(i+1))
		files[i] = &File{id: i}
	}

	wg.Add(NUMS)

	for j := 0; j < NUMS; j++ {
		for k := 0; k < STEPS; k++ {
			file := files[j]
			file.contents = append(file.contents, <-chans[(j+k)%NUMS])
		}
		wg.Done()
	}

	for o := 0; o < NUMS; o++ {
		fmt.Println(files[o])
	}

}

func newChan(i int) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			ch <- i
		}
	}()
	return ch
}

type File struct {
	id       int
	contents []int
}

func (f *File) append(num int) {
	f.contents = append(f.contents, num)
}

func (f *File) String() string {
	return fmt.Sprintf("%s:%v", string(rune(65+f.id)), f.contents)
}
