package main

import (
	"fmt"
)

func main() {
	data := [5]int{1, 2, 3, 4, 5}
	s := data[:0]
	fmt.Println(s)
}
