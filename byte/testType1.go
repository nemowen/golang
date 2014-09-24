package main

import (
	"fmt"
)

type Integer int

func (this Integer) Add(p Integer) Integer {
	return this + p
}

func (this Integer) less(p Integer) bool {
	return this < p
}

func main() {
	var b Integer = 2

	fmt.Println(b.Add(3))

	fmt.Println(b.less(4))

}
