package main

import (
	"fmt"
	"reflect"
)

func foo() {
	println("foo")
}

func bar(a ...int) int {
	var sum int = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func ok(msg string) string {
	println("ok func")
	return "ok" + msg
}

func main() {
	funcs := make(map[string]interface{}, 5)
	funcs["foo"] = foo
	funcs["ok"] = ok
	funcs["bar"] = bar

	v2, _ := Call(funcs, "bar", 1, 2, 3, 4)
	i, v, ok := reflect.TypeOf(v2)
	fmt.Println(i, v, ok)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	vt := reflect.ValueOf(m[name])

	in := make([]reflect.Value, len(params))
	for i, v := range params {
		in[i] = reflect.ValueOf(v)
	}
	return vt.Call(in), nil

}
