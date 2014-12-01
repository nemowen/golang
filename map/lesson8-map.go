package main

import "fmt"

func main() {
	sm := make([]map[int]string, 5)
	for i, _ := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "OK"
		sm[i][2] = "error"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}
