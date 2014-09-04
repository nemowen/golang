package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	c := sync.NewCond(&m)
	n := 2
	running := make(chan bool, n)
	awake := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func() {
			m.Lock()
			running <- true
			c.Wait()
			awake <- true
			m.Unlock()
		}()
	}
	for i := 0; i < n; i++ {
		fmt.Println(i)
		<-running // Wait for everyone to run.
	}
	for n > 0 {
		select {
		case <-awake:
			fmt.Println("goroutine not asleep")
		default:
		}
		m.Lock()
		c.Signal()
		m.Unlock()
		<-awake // Will deadlock if no goroutine wakes up
		select {
		case <-awake:
			fmt.Println("too many goroutines awake")
		default:
		}
		n--
	}
	c.Signal()

}
