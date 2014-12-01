package main

import (
	"fmt"
	//"sync/atomic"
	"time"
)

func main() {
	start := time.Now()
	var ints = [10]int{14, 9, 6, 7, 8, 43, 69, 7, 53, 3}
	b := quick(ints, 0, len(ints))
	fmt.Println(b)
	fmt.Println(time.Now().Sub(start))
}

func quick(ints [10]int, low, high int) [10]int {

	var temp int
	if low < high {
		var i, j, key int = low, high, ints[low]

		for ; j > i; j-- {
			if ints[j] < key {
				temp = ints[i]
				ints[i] = ints[j]
				ints[j] = temp
			}
		}

		for ; i < j; i++ {
			if ints[i] > key {
				temp = ints[j]
				ints[j] = ints[i]
				ints[i] = temp
			}
		}

		ints[i] = key
		quick(ints, low, i-1)
		quick(ints, i+1, high)
	}
	return ints
}
