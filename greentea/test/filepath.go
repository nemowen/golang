package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan int)
	t, _ := time.Parse("2006-01-02 15:04:05", "2014-05-20 17:22:00")
	//println(t.Add(1 * time.Hour).String())
	fmt.Println(time.Now().Sub(t.Add(2*24*time.Hour)) > 0)

	tick := time.Tick(2 * time.Second)

	go func() {
		for {
			select {
			case <-tick:
				fmt.Println((time.Since(t) + (8 * time.Hour)).Seconds())
			}
		}
	}()
	<-exit
}
