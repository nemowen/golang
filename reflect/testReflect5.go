package main

import (
	"reflect"
)

func main() {
	s := reflect.ValueOf([]int{1, 2, 3, 4})
	for i, n := 0, s.Len(); i < n; i++ {
		println("slince:", s.Index(i).Int())
	}

	mm := map[string]string{"age": "27", "name": "Nemo", "sex": "man"}
	m := reflect.ValueOf(&mm).Elem()
	println("map len:", m.Len())
	println("map canset:", m.CanSet())

	m.SetMapIndex(reflect.ValueOf("name"), reflect.ValueOf("Nora"))
	for i, v := range m.MapKeys() {
		println(i, v.String(), m.MapIndex(v).String())
	}

}
