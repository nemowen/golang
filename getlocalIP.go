package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		ip := addr.String()
		if "0.0.0.0" != ip {
			fmt.Println(ip)
		}
	}
}
