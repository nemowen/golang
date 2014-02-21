package main

import "fmt"
import "os"

func sayHi(msg string, c chan int) {

	for i := 0; i < 5; i++ {
		//runtime.Gosched()
		if "word" == msg {
			fmt.Println(<-c)
		} else {
			c <- i
		}
		fmt.Println(msg)
	}
}

func main() {
	ch := make(chan int, 5)
	fmt.Println(cap(ch))
	go sayHi("word", ch)
	sayHi("hello", ch)

	s, e := os.Hostname()
	if e != nil {
		fmt.Println(s)
	} else {
		fmt.Println(e)
	}
}
