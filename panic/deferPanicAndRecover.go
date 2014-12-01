package main

import "fmt"

func main() {
	var i = 5
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("error happend in main... %s\n", r)
		}
		fmt.Println("no error")
	}()

	f(i)
	var a = &i
	fmt.Println(*a)
}

func f(i int) {
	for {
		if i > 3 {
			panic("nemo")
		}
		fmt.Printf("number is %d\n", i)
		i += 1
	}
}
