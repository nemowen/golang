package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type user struct {
	id   int
	name string
	c    complex128
}

func main() {
	u := user{1, "Nemo", 2}
	fmt.Println(unsafe.Sizeof(u.c))
	var a float64 = 1 << 100
	fmt.Printf("%v\n", a)
	fmt.Printf("System Name:%s, System Arch:%s, CPU Numbers:%d", runtime.GOOS, runtime.GOARCH, runtime.NumCPU())

}
