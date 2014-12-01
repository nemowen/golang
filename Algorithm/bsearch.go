package main

import (
	"fmt"
)

/**
折半查询
**/

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%d\n", bsearch(a, 7))
	fmt.Printf("%d\n", bsearch_rec(a, 0, len(a), 7))
}

func bsearch(a []int, key int) int {
	var low, high int = 0, len(a)
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if key == a[mid] {
			return mid
		}
		if key < a[mid] {
			high = mid - 1
		} else if key > a[mid] {
			low = mid + 1
		}
	}
	return -1
}

/**
使用了递归算法
**/
func bsearch_rec(a []int, low, high, key int) int {
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if key == a[mid] {
			return mid
		}
		if key < a[mid] {
			return bsearch_rec(a, low, mid-1, key)
		}
		if key > a[mid] {
			return bsearch_rec(a, mid+1, high, key)
		}
	}
	return -1
}
