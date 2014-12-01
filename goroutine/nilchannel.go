package main

import (
	"fmt"
	"net"
	"os"
	//"time"
)

func waitMany(a, b chan bool) {
	for a != nil || b != nil {
		select {
		case <-a:
			a = nil
			fmt.Println("case a")
		case <-b:
			b = nil
			fmt.Println("case b")
		}
	}
}

func main() {

}
