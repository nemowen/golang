package main

<<<<<<< HEAD
type callback func(s string)

func test(a, b int, sum func(int, int) int) {
	println(sum(a, b))
}

func main() {
	var cb callback
	cb = func(s string) {
		println(s)
	}
	cb("hello, world!")

	sum := func(a, b int) int {
		return a + b
	}

	test(1, 2, sum)
}
=======
import "fmt"
>>>>>>> f3e35a66f80e382f86e6bf5e04031b9ae272304a
