package main

import (
	"fmt"
	"reflect"
)

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func main() {
	c := reflect.ChanOf(reflect.SendDir|reflect.RecvDir, Int)
	fmt.Println(c)

	t := reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(t)

}
