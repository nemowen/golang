package main

import (
	"fmt"
	"time"
)

func main() {
	ticket := time.Tick(time.Millisecond * 1000) //time.NewTicker(time.Second * 1)
	for {
		select {
		case i := <-ticket:
			fmt.Println(i.Format("2006-01-02 15:04:05"))
		}
	}

}
