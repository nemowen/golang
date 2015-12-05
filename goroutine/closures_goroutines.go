package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	for ix := range values {
		func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()

	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println()

	for ix := range values {
		go func(ix int) {
			fmt.Print(ix, " ")
		}(ix)
	}
	time.Sleep(time.Second * 2)
	fmt.Println()

	for ix := range values {
		v := values[ix]
		go func() {
			fmt.Print(v, " ")
		}()
	}
	time.Sleep(time.Second * 2)
}
