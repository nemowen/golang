package main

import (
	"reflect"
)

type Date struct {
	a bool
	b int64
	c string
}

func main() {
	var d Date

	t := reflect.TypeOf(d)
	println(t.Size(), t.Align())

	a, _ := t.FieldByName("c")
	println(a.Type.Size(), a.Type.Align())
}
