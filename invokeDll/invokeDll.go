package main

import (
	"fmt"
	"syscall"
)

var versionAddr uintptr

func init() {
	k32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		panic("LoadLibrary " + err.Error())
	}
	defer syscall.FreeLibrary(k32)
	versionAddr, err = syscall.GetProcAddress(k32, "GetVersion")
}

func main() {
	r1, r2, _ := syscall.Syscall(versionAddr, 0, 0, 0, 0)
	fmt.Printf("%s,%s", r1, r2)
}
