package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	handwareAddrs map[string]string
)

func main() {
	// 以下读取网卡信息
	Interface, err := net.Interfaces()
	if err != nil {
		panic("未发现网卡地址")
		os.Exit(1)
	}

	handwareAddrs = make(map[string]string, len(Interface))
	for _, inter := range Interface {
		inMAC := strings.ToUpper(inter.HardwareAddr.String())
		handwareAddrs[inMAC] = inMAC
	}

	if len(os.Args) != 2 {
		fmt.Println("为保障安全:请先绑定本机上的网卡地址")
		os.Exit(0)
	}

	addr := os.Args[1]
	h, e := net.ParseMAC(addr)
	if e != nil {
		fmt.Println("为保障安全:请先绑定本机上的网卡地址")
		fmt.Println("方法：client.exe 90-4C-E5-58-7E-FE")
		os.Exit(2)
	}
	inputMAC := strings.ToUpper(h.String())
	if inputMAC != handwareAddrs[inputMAC] {
		fmt.Println("网卡地址不匹配")
		os.Exit(0)
	}
}
