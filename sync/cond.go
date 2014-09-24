//Cond 在 Locker 的基础上增加了一种 "通知" 机制。其内部通过一个计数器和信号量来实现通知和广播的效果。

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var cpu int = 1
	runtime.GOMAXPROCS(cpu)

	var wg = new(sync.WaitGroup)
	var cond = sync.NewCond(new(sync.Mutex))

	// 生成两个Goroutine执行
	for id := 0; id < 3; id++ {
		wg.Add(1)
		go func(id int) {
			// 调用wait方法让每个Goroutine等待
			fmt.Println("call cond.Wait()")
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()

			// 每个Goroutine处理的事情
			for i := 0; i < 100; i++ {
				fmt.Println(id, ":->", i)
			}

			wg.Done()
		}(id)
	}

	fmt.Println("休息2秒钟,等待所有goroutine处于阻塞状态。（关键）")
	time.Sleep(time.Second * 2)

	fmt.Println("开始广播，通知所有goroutine的等待停止，开始执行")
	cond.L.Lock()
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()
	fmt.Println("main end")

}
