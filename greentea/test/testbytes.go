package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := []byte{'e', 'a', 'b', 'a', 'b', 'd'}
	si := bytes.Index(buffer, []byte("ea"))
	ei := bytes.Index(buffer, []byte("bd"))
	fmt.Printf("%s", buffer[si+2:ei])
}
