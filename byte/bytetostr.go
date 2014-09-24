package main

import (
	"fmt"
)

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12}
	a := data[:3]
	fmt.Println(a, len(a), cap(a))
}
