package main

import "fmt"

func main() {

	s := [5]int{1, 2, 3, 4, 5}

	s1 := s[0:3]
	sum := make([]int, 0, 10)

	//  s1和s2指向同一个底层数组，copy只是数据上的变化，而没有影响到各个切片的指向位置！
	//copy(s2, s1)

	//sum = append(sum, s1...)
	//sum = append(sum, s[:]...)
	fmt.Printf("%p %d %d\n", sum, len(sum), cap(sum))
	sum = append(sum, s1...)
	sum = append(sum, 4, 5, 6, 7, 8, 9, 10, 11)

	fmt.Println(sum)
	fmt.Printf("%p %d %d\n", sum, len(sum), cap(sum))

}
