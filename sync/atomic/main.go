package main

/**
	原子操作测试:
	原子操作(atomic operation)是不需要synchronized
**/

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i int64 = 1

	atomic.AddInt64(&i, 5) // 原子操作加5
	fmt.Println(i)

	b := atomic.CompareAndSwapInt64(&i, 6, 10) // 如果第一个参数，与第二个参数相等，则改变值为3
	fmt.Println(i, b)

	atomic.StoreInt64(&i, 5) // 付值操作
	fmt.Println(i)

}
