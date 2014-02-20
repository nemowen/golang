package main

import (
	"fmt"
)

//借助于接口，Go 完全可以实现 OOP 所需的操作。

type User struct {
	id   int
	name string
}

type Manager struct {
	group string
	User
}

type Tester interface {
	test()
}

func (this User) test() {
	fmt.Println(this)
}

func (this Manager) test() {
	fmt.Println(this)
}

func dosomething(o Tester) {
	o.test()
}

func main() {
	m := Manager{"IT", User{1, "Nemo"}}
	dosomething(m)
	dosomething(m.User)
}
