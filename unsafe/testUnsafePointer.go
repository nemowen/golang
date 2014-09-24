package main

import (
	"fmt"
	//"text/template"
	"unsafe"
)

// MyTemplate 定义和 template.Template 只是形似
type MyTemplate struct {
	name       string
	parseTree  *unsafe.Pointer
	common     *unsafe.Pointer
	leftDelim  string
	rightDelim string
}

func main() {
	u := uint32(32)
	i := int32(1)
	//println(&u, &i)
	p := &i
	//p = &u
	//p = (*int32)(&u)
	p = (*int32)(unsafe.Pointer(&u))
	fmt.Println(p)

	t := new(interface{})
	fmt.Println(t)
	p2 := (*MyTemplate)(unsafe.Pointer(&t))
	p2.name = "Wen"
	p2.leftDelim = "re"
	p2.rightDelim = "23"
	fmt.Println(p2, t)
}
