// 读写锁 RWMutex

package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(2)

	// 创建读写锁对象
	var rw = new(sync.RWMutex)
	var wg = new(sync.WaitGroup)

	for id := 0; id < 3; id++ {
		wg.Add(1)
		go func(name string) {
			if name == "goroutine1" {
				time.Sleep(1 * time.Second)
			}

			// 模拟完成后等待，给时间给写操作
			rw.RLock()
			for i := 0; i < 10; i++ {
				fmt.Println(name, ":", i)
			}
			// goroutine任务完成
			wg.Done()
			rw.RUnlock()

		}("goroutine" + strconv.Itoa(id))
	}

	time.Sleep(1 * time.Second)

	rw.Lock()
	fmt.Println("进入写操作。。。")
	time.Sleep(5 * time.Second)
	fmt.Println("写操作完成。。。")
	rw.Unlock()

	wg.Wait()
}
