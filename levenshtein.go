package main

import "fmt"

//import s "strings"

/*
func compare(str string, target string) int {
	var i, j, temp int
	var ch1, ch2 string
	n := len(str)
	m := len(target)

	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	d := make([n + 1][m + 1]int)

	for i = 0; i <= n; i++ { // 初始化第一列
		d[i][0] = i
	}

	for j = 0; j <= m; j++ { // 初始化第一行
		d[0][j] = j
	}

	for i = 1; i <= n; i++ {
		//ch1=

	}

}*/

func main() {
	s := "中国人asd的ee"
	for i, c := range s {
		fmt.Printf("%d = %c\n", i, c)
	}

}
