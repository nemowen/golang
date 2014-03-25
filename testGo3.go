package main

import (
	"fmt"
)

type Interface interface {
	say() string
}

type Object struct {
	o int
}

func (this *Object) say() string {
	this.o++
	return "hello"
}

func do(i Interface) string {
	return i.say()
}

func main() {
	o := &Object{}
	fmt.Println(do(o))
	fmt.Printf("Object Type:%T \n", o)
	fmt.Printf("Object o:%d \n", o.o)
}
