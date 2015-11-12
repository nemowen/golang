package main

import (
	"fmt"
	"reflect"
)

// slice ,map 操作
func main() {
	s := make([]int, 0, 10)
	sv := reflect.ValueOf(&s).Elem()
	sv.SetLen(2)
	sv.Index(0).SetInt(1)
	sv.Index(1).SetInt(2)
	sv = reflect.Append(sv, reflect.ValueOf(200))

	sv = reflect.AppendSlice(sv, reflect.ValueOf([]int{300, 400}))
	fmt.Println(sv.Interface())

	m := map[string]int{"a": 1}
	sv = reflect.ValueOf(&m).Elem()
	sv.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) // update
	sv.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) // add
	fmt.Println(sv.Interface())
}
