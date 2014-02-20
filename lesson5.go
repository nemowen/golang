package main

import "fmt"

const (
	B float64 = (1 << (iota * 10))
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type a struct {
	a int `PK` //`PK` 什么意思，有什么用？
}

func main() {

	//fmt.Println(YB)
	// if a := 1; a > 1 {
	// 	fmt.Println(a)
	// } else if b := 2; b > 1 {
	// 	fmt.Println(a)
	// }

	aa := []int{1, 2, 3, 4, 5}

	s1 := aa[2:5]

	s2 := aa[1:3]
	fmt.Println(s1, s2, aa)

	s1 = append(s1, 5, 6, 7, 8, 9, 0, 1, 1, 1)

	s1[1] = 1000

	fmt.Println(s1, s2, aa)
}
