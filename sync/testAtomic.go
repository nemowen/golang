package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a int32 = 10
	atomic.AddInt32(&a, 63)
	b := atomic.SwapInt32(&a, 44)
	fmt.Println(b)
	fmt.Println(a)
	if atomic.CompareAndSwapInt32(&a, 10, 12) {
		fmt.Println(a)
		fmt.Println("ok")
	} else {
		fmt.Println("err")
	}
}
