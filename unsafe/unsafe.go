package main

import (
	"unsafe"
)

type data struct {
	name string
	age  int
}

func main() {
	d := data{name: "Nemo", age: 28}
	println(unsafe.Pointer(&d))
	p := unsafe.Pointer(&d)
	s := (*string)(p)
	println(*s, s)

	u := uintptr(p)
	u += unsafe.Offsetof(d.name)

	ci := (*string)(unsafe.Pointer(u))
	*ci = *ci + " Wen"

	println(d.name, &d.name)

}
