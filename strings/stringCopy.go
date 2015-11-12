package main

import (
	"fmt"
)

func main() {
	str := "A string"
	bytes := make([]byte, len(str))
	i := copy(bytes, str)
	copystr := string(bytes)
	fmt.Printf("%x %d\n", &str, i)
	fmt.Printf("%x\n", &copystr)
}
