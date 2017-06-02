package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {
	cond.L.Lock() //获取锁
	cond.Wait()   //等待通知  暂时阻塞
	fmt.Println(x)
	time.Sleep(time.Second * 1)
	cond.L.Unlock() //释放锁
}
func main() {
	runtime.GOMAXPROCS(4)
	for i := 0; i < 10; i++ {
		go test(i)
	}

	fmt.Println("start all")
	time.Sleep(time.Second * 3)
	fmt.Println("broadcast one")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine

	time.Sleep(time.Second * 3)
	fmt.Println("broadcast one")
	cond.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine

	time.Sleep(time.Second * 3)
	fmt.Println("broadcast all")
	cond.Broadcast() //3秒之后 下发广播给所有等待的goroutine

	time.Sleep(time.Second * 10)

}
