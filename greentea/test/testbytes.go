package main

import (
	//"bytes"
	"fmt"
	//"runtime"
	//"os"
	//"strconv"
	//"strings"
)

func main() {
	buffer := []byte{'e', 'a', 'b', 'a', 'b', 'd'}
	fmt.Printf(">>%p\n", buffer)
	buffer[1] = 'e'
	fmt.Printf(">>%p\n", buffer)
}
