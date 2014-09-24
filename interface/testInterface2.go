package main

/**
	把接口类型匿名嵌入到 struct
**/

import (
	"fmt"
)

// define a interface
type ITest interface {
	test()
}

// define a struct
type Mytest struct {
}

// define an User struct
type User struct {
	id   int
	name string
	ITest
}

// define a method of Mytest
func (Mytest) test() {
	fmt.Println("test ok")
}

func main() {
	u := User{1, "Nemo", nil}
	u.ITest = new(Mytest)
	fmt.Println(u)
	u.test()
}
