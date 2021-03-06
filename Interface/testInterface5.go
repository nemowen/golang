package main

import "fmt"

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

//
func (*User) test() {
	fmt.Println("user test")
}

func main() {
	u := User{1, "Nemo"}
	m := Manager{"管理者", User{2, "Nora"}}

	var it Tester
	it = &u
	it.test()

	it = &m
	it.test()
}
