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

	var ia interface{}
	fmt.Println(unsafe.Sizeof(ia), ia == nil)

}
