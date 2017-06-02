package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	for k := 1; k <= 3; k++ {
		x = 4 * k
		y = 25 - 7*k
		z = 75 + 3*k
		fmt.Printf("公鸡:%d只,母鸡:%d只,小鸡:%d只\n", x, y, z)
	}
}
