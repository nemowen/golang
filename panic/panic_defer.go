package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("returned normally from f.")
}

func f() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered in f", err)
		}
	}()
	fmt.Println("caling g.")
	g(0) // when panic, the program is stoped, but the defer is ran normally
	fmt.Println("returned normally from g.")

}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in g:", i)
	fmt.Println("printing in g ", i)
	g(i + 1)
}

//caling g.
//printing in g 0
//printing in g 1
//printing in g 2
//printing in g 3
//Panicking
//defer in g:3
//defer in g:2
//defer in g:1
//defer in g:0
//recovered in f 4
//returned normally from f.
