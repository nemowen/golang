package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
	Job  string
}

func editObj(i interface{}) {
	v := reflect.ValueOf(i)

	// 首先判断是 Ptr 类型
	// 用 Elem() 获取 ptr 指向的实际对象，并判断是 settable。
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		fmt.Println("Obj cannot Setable!")
		return
	} else {
		v = v.Elem() //我们要操作的是目标对象，而不是ptr
	}

	field := v.FieldByName("Name")
	if field.Kind() == reflect.String {
		field.SetString("Jack")
	}

}

func main() {
	u := User{1, "Nemo", 27, "Programmer"}
	editObj(&u)
	fmt.Println(u)
}
