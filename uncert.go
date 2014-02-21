package main

import "fmt"

func main() {
	uncert(1, 3, 4, 5, 6, 7)
}

func uncert(args ...int) {
	fmt.Printf("%T , [%d]", args, len(args))
}
