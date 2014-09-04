package main

import (
	//"fmt"
	"strconv"
)

func main() {
	var a int64
	a = 0
	a = a | (1 << 3)  // 在 bit3 上设置标志位 (从 bit0 开始算)
	a = a | (1 << 6)  // 在 bit6 上设置标志位
	a = a &^ (1 << 6) // 清除 bit6 上的标志位，a = 8 = 0000 0100
	println(strconv.FormatInt(a, 2))

	// Convert string witch represent binary number into int
	// 二进制转十进制
	if i, err := strconv.ParseUint("111100", 2, 64); err != nil {
		println(err)
	} else {
		println(i)
	}

	// Converting from an integer to its binary representation
	// 十进制转二进制
	println(strconv.FormatInt(128, 2))

	//println(strconv.FormatInt("11111111", 2, 64))

}
