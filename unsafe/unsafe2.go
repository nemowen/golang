package main

import "fmt"
import "unsafe"
import "reflect"

type D struct {
	a int
	b string
}

func main() {
	//var p int64 = 23
	d := D{a: 10, b: "ab"}
	t := reflect.ValueOf(d)
	fmt.Println(t.FieldByName("a")) // use reflect package get field's value of D struct
	//size := unsafe.Sizeof(p)
	//fmt.Println(size)
	//fmt.Println(p)
	pit := (*int)(unsafe.Pointer(&d.a))
	*pit += 4
	fmt.Println(d)
}
