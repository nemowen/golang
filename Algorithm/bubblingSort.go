package main

import (
	"fmt"
	"time"
)

/**
冒泡排序
**/
func main() {
	start := time.Now()
	var ints = [10]int{69, 9, 6, 7, 8, 43, 14, 7, 53, 3}
	var sum int
	var temp int
	var lens = len(ints)
	for i := 0; i < lens; i++ {
		for j := 0; j < lens-1; j++ {
			if ints[j] > ints[j+1] {
				temp = ints[j]
				ints[j] = ints[j+1]
				ints[j+1] = temp
			}
		}
	}
	fmt.Println(ints, sum)
	fmt.Println(time.Now().Sub(start))
}
