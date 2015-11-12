package main

import (
	"fmt"
	"reflect"
)

type Date struct {
}

func (d Date) String() string {
	return "12rweer"
}

func main() {
	var d Date
	to(&d)

	t := reflect.TypeOf(d)
	println("Date:", t.NumMethod())

	a := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	println("stringer:", t.NumMethod())
	fmt.Println(t.Implements(a))
}

func to(i interface{}) {
	if a, o := i.(fmt.Stringer); o {
		fmt.Println(a.String())
	}
}
