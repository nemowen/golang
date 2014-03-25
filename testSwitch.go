package main

import (
	"fmt"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough // 直接运行case 3 下的语句块，不进行条件判断
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("default")
	}
}
