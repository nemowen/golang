package main

import (
	"fmt"
)

func main() {
	var sum int = 0
RESTART:
	fmt.Println("start for ... ")
	for i := 0; i < 10; i++ {
		if i == 5 {
			sum++
			if sum < 5 {
				goto RESTART
			} else {
				goto END
			}

		}
		fmt.Println("num:", i)
	}
END:
	fmt.Println("end run ... ")
}
