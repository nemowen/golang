package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ", os.Args[0], "ip:port")
		os.Exit(1)
	}

	addr := os.Args[1]
	fmt.Println(addr)

	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		fmt.Println(inter.Name, inter.HardwareAddr)
	}
}
