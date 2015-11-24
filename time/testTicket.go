package main

import (
	"fmt"
	"time"
)

func main() {
	ticket := time.Tick(time.Second * 1) //time.NewTicker(time.Second * 3)
	//	for {
	//		select {
	//		case <-ticket:
	//			println("ok")
	//		}
	//	}
	for i := range ticket {
		fmt.Println("ok", i)
	}
	//defer ticket.Stop()
	//time.t

}
