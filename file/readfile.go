package main

import (
	"fmt"
	"os"
)

// if the data columns are separate by a space, you can use the FScan-function
// series from the "fmt" package.
//
// data simple:
// name age state
// nemo 28	ok
// nora 27 ok
//
// returned
// [name nemo nora]
// [age 28 27]
// [state ok ok]
func main() {
	f, e := os.Open("C:/Users/nemowen/testfscan.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	var c1, c2, c3 []string
	for {
		var v1, v2, v3 string
		_, e := fmt.Fscanln(f, &v1, &v2, &v3)

		if e != nil {
			break
		}
		c1 = append(c1, v1)
		c2 = append(c2, v2)
		c3 = append(c3, v3)
	}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}
