package main

import ()

func main() {
	for i := 0; i < 3; i++ {
		if i == 2 {
			break
		}
	}
}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
