package main

import (
	"fmt"
)

type A struct {
}

type B struct {
}

type I interface {
	test()
}

func (a *A) test() {
	fmt.Println("test a")
}
func (b *B) test() {
	fmt.Println("test b")
}

func typeOf(i interface{}) {
	switch t := i.(type) {
	case *A:
		fmt.Printf("A Type %T with value %v\n", t, t)
		break
	case B:
		fmt.Printf("B Type %T with value %v\n", t, t)
		break
	case *B:
		fmt.Printf("*B Type %T with value %v\n", t, t)
		break
	default:
		fmt.Printf("default Type %T with value %v\n", t, t)
		break
	}
}

func ooo(i I) {
	i.test()
}

func main() {
	a := new(A)
	b := B{}
	a.test()
	b.test()

	ooo(a)
	ooo(&b)
}
