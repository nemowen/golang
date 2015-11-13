package main

import (
	"fmt"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "65.12 / 5212 / GO"
	fomat                  = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi firstName:%s lastName:%s!\n", firstName, lastName)
	fmt.Sscanf(input, fomat, &f, &i, &s)
	fmt.Println(f, i, s)

}
