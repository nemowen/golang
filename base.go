package main

import "fmt"

func main() {
	a := "test1"
	for _, v := range a {
		d := string(v)
		fmt.Println(d)
	}
}
