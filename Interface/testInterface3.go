package main

import (
	"fmt"
	"unsafe"
)

// define a interface
type ITest interface {
	test()
}

// define an User struct
type User struct {
	id   int
	Name string
}

// define a method of Mytest
func (User) test() {
	fmt.Println("test ok")
}

func main() {
	u := User{1, "Nemo"}
	i := ITest(&u)
	fmt.Println(i.(*User).Name)

	fmt.Println(u, i)
	i.test()

	var ia interface{}
	var a int = 3289
	fmt.Println(unsafe.Sizeof(a), ia == nil)

}
