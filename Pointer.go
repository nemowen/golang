package main

import (
	"fmt"
)

func main() {
	var str = "Nemo"
	gotos1(str)
	fmt.Println(str)
	gotos2(&str)
	fmt.Println(str)
}

func gotos1(str string) {
	str = "Nora"
}

func gotos2(str *string) {
	*str = "Nora"
}
