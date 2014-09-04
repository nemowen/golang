package main

import (
	"fmt"
)

type Tester interface {
	test()
}

type Stringer interface {
	tostring()
}

type User struct {
	id   int
	name string
}

func (User) test() {
	fmt.Println("user test method")
}

func (*User) tostring() {
	fmt.Println("user tostring method")
}

func main() {
	u := User{1, "Nemo"}
	p := &u

	Tester(u).test()
	Tester(p).test()

	//Stringer(u).tostring()
	Stringer(p).tostring()
}
