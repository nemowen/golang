package main

import (
	"fmt"
	"log"
	"os"
)

type AF func()

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func test() {
	f, err := os.Open("c:/ok.txt")
	CheckError(err)
	defer f.Close()
}

func HandlerError(a AF) AF {
	return func() {
		defer func() {
			if e, ok := recover().(error); ok {
				log.Printf("runtime panic: %v", e)
			}
		}()
		a()
	}
}

func main() {
	f := HandlerError(test)
	fmt.Println("................")
	f()
}
