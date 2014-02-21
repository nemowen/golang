package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	lenr := len(runes)
	fmt.Println(lenr)
	for i, j := 0, lenr-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	a := Reverse("wo you ok")
	b := strings.Title(a)
	fmt.Println(b)
}
