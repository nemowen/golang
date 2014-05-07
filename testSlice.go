package main

import (
	"fmt"
)

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	ms := make([]int, 2)
	ms[0] = 0
	ms[1] = 1
	// cap 每一次增加一倍
	fmt.Printf("%p,%v,%d\n", ms, ms, cap(ms))
	ms = append(ms, 3)
	fmt.Printf("%p,%v,%d\n", ms, ms, cap(ms))
	ms = append(ms, 4)
	fmt.Printf("%p,%v,%d\n", ms, ms, cap(ms))
	ms = append(ms, 5)
	fmt.Printf("%p,%v,%d\n", ms, ms, cap(ms))

	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[2:5]
	fmt.Println(s1, "len:", len(s1), "cap:", cap(s1))

	s2 := s1[2:6:7]
	fmt.Println(s2, "len:", len(s2), "cap:", cap(s2))

	s3 := s2[3:5]
	fmt.Println(s3, "len:", len(s3), "cap:", cap(s3))
}
