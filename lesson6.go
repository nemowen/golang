package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[0:]
	s2[0] = 10
	fmt.Println(s1)
	fmt.Println(s2)
}
