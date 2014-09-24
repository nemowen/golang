/**
NumGoruntine init os.agrs 的一些例子
*/
package main

import (
	"os"
	"runtime"
	"time"
)

func init() {
	println("init1")
}

func init() {
	time.Sleep(2 * time.Second)
	println("init2")
}

func main() {
	// os.Exit(0) 返回终⽌止进程
	for _, v := range os.Args { // os.Args 获取命令行启动参数。
		println("args:", v)
	}
	runtime.GOMAXPROCS(1)

	println("start:", runtime.NumGoroutine()) // NumGoroutime : 等待排队的Goroutine数量

	for i := 0; i < 10; i++ {
		go func(n int) {
			println(n, runtime.NumGoroutine())
		}(i)
	}

	time.Sleep(3 * time.Second)
	println("over:", runtime.NumGoroutine())
}
