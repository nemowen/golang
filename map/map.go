package main

import (
	"fmt"
	"math/cmplx"
	"strconv"
)

type Maps struct {
	m, b float32
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 4 >> 1
	z      complex128 = cmplx.Sqrt(4)
)

func main() {
	/*q := make(map[string]Maps)
	fmt.Println(q)
	q["num"] = Maps{23.12, 43.34}

	q["count"] = Maps{45.23, 657.34}
	delete(q, "num")

	a, b := q["count"]

	fmt.Println(a, "bb", b)
	*/
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	var a = 65
	//b1:=string(a)  // b1=A
	b := strconv.Itoa(a) //将int 转成文本
	fmt.Println(b)

}
