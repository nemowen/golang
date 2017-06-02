package main

import "fmt"

func decoration(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("started")
		f(s)
		fmt.Println("done")
	}
}

func Hello(s string) {
	fmt.Println("doing", s)
}

func main() {
	hello := decoration(Hello)
	hello("lala")
}
