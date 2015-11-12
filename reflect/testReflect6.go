// 指针操作非导出字段
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Name string
	age  int
}

func main() {
	u := User{Name: "nemo", age: 27}
	v := reflect.ValueOf(&u).Elem()
	v.FieldByName("Name").SetString("Nora")

	a := v.FieldByName("age")
	fmt.Println("is age canSet:", a.CanSet())
	if a.CanAddr() {
		age := (*int)(unsafe.Pointer(a.UnsafeAddr()))
		*age = 2
	}
	fmt.Println(v.Interface().(User))
}
