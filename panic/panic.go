package main

import "fmt"

/**********************
	run and result:

	It's A
	the error message is : func b is error
	It's C
**********************/

func main() {

	echoA()
	echoB()
	echoC()

}

func echoA() {
	fmt.Println("It's A")
}

func echoB() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("the error message is : %s \n", err)
		}
	}()

	panic("func b is error")

	fmt.Println("It's B")
}

func echoC() {
	fmt.Println("It's C")
}
