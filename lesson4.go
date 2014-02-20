package main

import (
	"fmt"
)

const (
	A = "Q"
	B = 2
	C = iota
	D
	F
)

//1110

//位运算
// 	6 : 0110
// 11 : 1011
//-----------
// &  : 0100
// |  : 1111
// ^  : 1011
// &^ : 0100

func main() {
	// fmt.Println(A)
	// fmt.Println(B)
	// fmt.Println(C)
	// fmt.Println(D)
	// fmt.Println(F)

	fmt.Println(1 &^ 15)

}
