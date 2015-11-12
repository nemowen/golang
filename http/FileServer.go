package main

import (
	"fmt"
	//"math/rand"
	//"time"
	//"strconv"
)

type T struct {
	initer Initer
	Name   string
}

type F struct {
	Name string
}

type A struct {
	T
}

func (t T) String() string {
	return "this T String() method"
}

func (t T) init() {
	t.initer.init()
	fmt.Println("t init")
}

func (a A) init() {
	fmt.Println("a init")
}

type Initer interface {
	init()
}

func main() {
	t := T{initer: new(A), Name: "2323"}
	t.init()

}
