package main

import (
	"fmt"
	"reflect"
)

// 用 reflection 机制写的一个比较通用的修饰器
func Decorator(decorationPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decorationPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(), func(args []reflect.Value) (results []reflect.Value) {
		fmt.Println("before")
		results = targetFunc.Call(args)
		fmt.Println("after")
		return
	})

	decoratedFunc.Set(v)
	return
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}

func main() {
	mybar := bar
	Decorator(&mybar, bar)
	mybar("hello", "world!")

}
