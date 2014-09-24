// 读写锁 RWMutex

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var rw = new(sync.RWMutex)

func Read(name string) {
	fmt.Println("read...")
	rw.RLock()
	fmt.Println("RLocking...")
	defer rw.RUnlock()

	for i := 0; i < 3; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	exit := make(chan bool, 1)
	fmt.Println("main")
	for i := 0; i < 3; i++ {
		go Read("goroutine" + strconv.Itoa(i))
	}
	time.Sleep(3 * time.Second)

	exit <- true
}
