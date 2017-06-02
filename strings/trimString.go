package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "\tthe important rôles of utf8 text\n"
	str = strings.Trim(str, "\t\r\n ")
	fmt.Printf("%s\n", str)

	a := strings.Contains("/fghgh//", "//")
	fmt.Println(a)

	b := strings.TrimPrefix("/aa/bb/", "/")
	fmt.Println(b)

	c := strings.TrimSuffix("/aa/bb/", "b/")
	fmt.Println(c)

}
