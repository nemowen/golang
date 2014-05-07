package main

import (
	"sync"
)

var c = make(chan int)
var once sync.Once

var a int = 0

func f() {
	once.Do(func() {
		a += 1
		c <- 0
	})
	println("ok")
}
func main() {

	go f()
	go f()
	go f()
	go f()
	<-c
	print(a)

}
