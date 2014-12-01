package main

import (
	"fmt"
	"time"
)

/**
选择排序
**/
func main() {
	start := time.Now()
	var ints = [10]int{14, 9, 6, 7, 8, 43, 69, 7, 53, 3}
	var lens = len(ints)
	var j, min, tmp int
	for i := 0; i < lens; i++ {
		min = i
		for j = i + 1; j < lens; j++ {
			if ints[j] < ints[min] {
				min = j
			}
			if min != i {
				tmp = ints[i]
				ints[i] = ints[min]
				ints[min] = tmp

			}
		}
	}
	fmt.Println(ints)
	fmt.Println(time.Now().Sub(start))
}
