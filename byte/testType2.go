package main

// 5.5 类型推断
// 利用类型推断，可以判断接口对象是否是某个具体的接口或类型。

import (
	"fmt"
)

type ITest interface {
	show()
}

type User struct {
	id   int
	name string
}

func (this *User) show() {
	fmt.Printf("%d: My Name is %s\n", this.id, this.name)
}

func Dosomething(i interface{}) {
	if o, ok := i.(ITest); ok {
		o.show()
	}

	i.(*User).show()
}

func main() {
	user := &User{1, "Nemo"}
	Dosomething(user)

}
