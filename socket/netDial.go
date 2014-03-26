package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
	}
	var b []byte

	n, err := conn.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println(b)
}
