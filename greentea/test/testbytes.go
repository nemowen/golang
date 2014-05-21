package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	buffer := []byte{'e', 'a', 'b', 'a', 'b', 'd'}
	fmt.Printf("%s %p\n", buffer, buffer)
	si := bytes.Index(buffer, []byte("ea"))
	ei := bytes.Index(buffer, []byte("bd"))
	fmt.Printf("%s\n", buffer[si+2:ei])
	fmt.Printf("count:%d\n", bytes.Count(buffer, []byte{'a'}))

	buffer = bytes.Replace(buffer, []byte("ea"), []byte("qq"), 1)
	fmt.Printf("%s %p\n", buffer, buffer)

	ch := "中国人"
	change := []rune(ch)
	change[1] = '美'
	fmt.Printf("%d %s\n", len(change), string(change))

	i, _ :=

		fmt.Println(i)
}
