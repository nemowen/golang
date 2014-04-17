package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(2))
	time.Sleep(1000)
	fmt.Println("hello word! 完完全！")

}
