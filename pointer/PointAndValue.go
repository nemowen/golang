package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b *B) change(d int) {
	b.thing = d
}

func (b B) wirte() string {
	return fmt.Sprintf("B.thing is : %d", b.thing)
}

func main() {
	var b B
	b.change(8)
	println(b.wirte())

	b1 := new(B)
	b1.change(9)
	println(b1.wirte())

	b2 := &B{}
	b2.change(10)
	println((*b2).wirte())
}
