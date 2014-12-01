package main

import (
	"fmt"
	"runtime"
	"sync"
)

type User struct {
	name string
	age  int
}

func main() {
	var pool sync.Pool

	var a = 1
	pool.Put(a)
	pool.Put(new(User))
	fmt.Println(pool.Get())
	runtime.GC()

	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
