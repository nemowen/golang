package main

import (
	"fmt"
)

func main() {
	c := make(chan []int)
	m := []int{1, 2, 3, 4}
	done := make(chan struct{})
	go func() {
		var i int
		for tt := range c {
			fmt.Println(i, ":", tt)
			i++
			if i == 4 {
				i = 0
			}
		}
	}()
	go func() {
		for i := 0; i < 4; i++ {
			n := i % 4
			mm := []int{}
			mm = append(mm, m[n:]...)
			mm = append(mm, m[:n]...)
			mm = append(mm, mm[:2]...)
			c <- mm
		}
		done <- struct{}{}
	}()
	<-done
}
