package main

import (
	"errors"
	"fmt"
)

func check1() (x int, err error) {
	return 1, errors.New("error for check1")
}
func check2(x int) (y int, err error) {
	return x + 1, errors.New("error for check2")
}

func shadow() (err error) {
	x, err := check1()
	if err != nil {
		//return err
	}
	if y, err := check2(x); err != nil {
		return err
	} else {
		fmt.Println(y)
	}
	return
}

func main() {
	err := shadow()
	fmt.Println(err)
}
