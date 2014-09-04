package main

import "fmt"

func main() {
	a, b := 11, 12
	fmt.Printf("a = %b\nb = %b\n", a, b)
	fmt.Printf("a & b = %b \n", a&b)
	fmt.Printf("a | b = %b \n", a|b)
	fmt.Printf("a ^ b = %b \n", a^b)
	fmt.Printf("a &^ b = %b \n", a&^b)

}
