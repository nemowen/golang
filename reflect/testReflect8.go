package main

import (
	"fmt"
	"reflect"
)

// reflect 调用方法执行
type Date struct {
}

func (Date) Swap(a, b int) (int, int) {
	return b, a
}

func (*Date) Sum(s string, b ...int) string {
	var sum int
	for _, v := range b {
		sum += v
	}
	return fmt.Sprintf(s, sum)
}

// 打印方法的信息
func info(m reflect.Method) {
	t := m.Type

	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Println("method in info:", t.Kind().String(), t.In(i), m.Name)
	}
	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Println("method out info:", t.Kind().String(), t.Out(i))
	}
}

func main() {
	var d Date
	t := reflect.TypeOf(d)
	m, b := t.MethodByName("Swap")
	if b {
		info(m)
	}
	m, b = t.MethodByName("Sum")
	if b {
		info(m)
	}

	fmt.Println("========================")

	f := reflect.ValueOf(&d)

	exec := func(method string, in []reflect.Value) {
		c := f.MethodByName(method)
		for _, v := range c.Call(in) {
			fmt.Println(method, "=", v.Interface())
		}
	}

	swapin := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}
	//_ = swapin
	sumin := []reflect.Value{
		reflect.ValueOf("result=%d"),
		reflect.ValueOf(5),
		reflect.ValueOf(5),
		reflect.ValueOf(5),
	}

	exec("Swap", swapin)
	exec("Sum", sumin)

}
