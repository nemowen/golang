package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

type Manager struct {
	Grout string
	User
}

type IUser interface {
	Test()
}

type IManager interface {
	Test()
	Test2()
}

func (this Manager) Test()  { fmt.Println(this) }
func (this Manager) Test2() { fmt.Println(this) }

func main() {
	var im IManager = Manager{"IT", User{1, "Nemo"}}
	im.Test()
	im.Test2()

	var iu IUser = im
	iu.Test()
}
