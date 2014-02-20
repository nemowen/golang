package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

func main() {
	//v := Vertex{1, 2}
	//v.x = 4
	//fmt.Println(v.x)

	//p := &v

	//p.x = 20

	//r := &p

	//s := &r

	//1). 结构体的声明可以简化；
	//2). 结构体有多种初始化；如果不指定具体的值，也都有默认值；
	//3). fmt.Printf 格式化输出，可以输出值的类型 %T %v %g；
	/*var (
		p = Vertex{1, 2}
		v = &Vertex{3, 4}
		a = Vertex{x: 5}
		b = Vertex{}
	)

	fmt.Println(p, *v, a, b)
	*/

	//1). new 构造结构体，返回指针，看来是借鉴了 c++ 的用法；
	//2). new 的同时貌似不能初始化，new 是个函数，不像 c++ 中是个运算符；
	v := new(Vertex)

	fmt.Println(*v, &v)

}
