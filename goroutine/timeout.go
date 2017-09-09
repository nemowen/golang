package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	setTimeout(ch)
	ch <- struct{}{}
	time.Sleep(1 * time.Second)
	ch <- struct{}{}
	time.Sleep(2 * time.Second)
	ch <- struct{}{}
	time.Sleep(5 * time.Second)
}

// the right way to set gorountine timeout
func setTimeout(ch chan struct{}) {
	go func() {
		timer := time.NewTimer(2 * time.Second)
		for {
			select {
			case v, _ := <-ch:
				fmt.Println(v)
			case <-timer.C:
				fmt.Println("Timeout...")
				timer.Reset(2 * time.Second)
			}

		}
	}()
}
