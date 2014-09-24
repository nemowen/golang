package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type Lesser interface {
	Less(b Integer) bool
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type IFly interface {
	Fly()
}

func main() {
	var a Integer = 23
	var b LessAdder = &a
	var c Lesser = &a
	b.Add(90)
	fmt.Println(a)

	if v, ok := c.(LessAdder); ok {
		fmt.Println(v.Less(24))
		v.Add(23)
		fmt.Println(a)
	} else {
		fmt.Println("not found.")
	}

}
