package main

import "fmt"

func sayHi(msg string, c chan int) {

	for i := 0; i < 5; i++ {
		//runtime.Gosched()
		if "word" == msg {
			fmt.Println(<-c)
		} else {
			c <- i
		}

	}
}

func main() {
	ch := make(chan int, 5)

	go sayHi("word", ch)
	sayHi("sasd", ch)

}
