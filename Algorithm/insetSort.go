package main

import (
	"fmt"
	"time"
)

/**
插入排序
**/
func main() {
	start := time.Now()
	var ints = [10]int{14, 9, 6, 7, 8, 43, 69, 7, 53, 3}
	var lens = len(ints)
	var temp int
	for i := 1; i < lens; i++ {
		for j := i; j >= 1; j-- {
			if ints[j-1] > ints[j] {
				temp = ints[j-1]
				ints[j-1] = ints[j]
				ints[j] = temp
			}
		}
	}
	fmt.Println(ints)
	fmt.Println(time.Now().Sub(start))
}
