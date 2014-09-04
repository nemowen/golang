package main

import (
	//"bytes"
	//"encoding/binary"
	"fmt"
	//"os"
	"strconv"
)

func main() {
	u, _ := strconv.ParseUint("111", 2, 10)
	fmt.Printf("%x\n", u)
	a := []byte("abcd")
	fmt.Printf("%b,%d", a, a)

}
