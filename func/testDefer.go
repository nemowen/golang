package main

import (
	"fmt"
	"time"
)

func main() {
	aa()
	for i := 10; i > 0; i-- {
		if i == 5 {
			defer func() {
				fmt.Println("i=", 5)
			}()
		}
		defer fmt.Println("i=", i)
		time.Sleep(2 * time.Second)
		fmt.Println(i)
	}
}

func aa() {
	defer func() {
		fmt.Println(">>>>>12")
	}()
	fmt.Println(">>>>>>>")
}
