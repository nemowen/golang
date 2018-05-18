package main

import (
	"fmt"
)

func main() {
	// to Uppercase
	var c = 'b'
	c -='a' - 'A'
	fmt.Println(string(c))

	// to Lowercase
	c = 'R'
	c += 'a'-'A'
	fmt.Println(string(c))
}
