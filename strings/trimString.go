package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "\tthe important rôles of utf8 text\n"
	str = strings.Trim(str, "\t\r\n ")
	fmt.Printf("%s\n", str)

}
