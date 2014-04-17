package main

import (
	"fmt"
)

// 常用位运算
func main() {
	fmt.Printf("去掉最后一位 | %b -> %b \n", 45, 45>>1)
	fmt.Printf("在最后一位加一个0 | %b -> %b \n", 45, 45<<1)
	fmt.Printf("把右数第5位变为1 | %b -> %b \n", 45, 45|(1<<(5-1)))
	fmt.Printf("把右数第4位变为0 | %b -> %b \n", 45, 45&^(1<<(4-1)))
	fmt.Printf("右数第3位取反 | %b -> %b \n", 45, 45^(1<<(3-1)))
	fmt.Printf("取末4位 | %b -> %b \n", 45, 45&((1<<4)-1))
	fmt.Printf("最右数4位 | %b -> %b \n", 45, 45>>(4-1)&1)
	fmt.Printf("把末4位全部变成1 | %b -> %b \n", 45, 45|((1<<4)-1))
	fmt.Printf("把末4位全部取反 | %b -> %b \n", 45, 45^((1<<4)-1))
	fmt.Printf("把右边连续的1变成0 | %b -> %b \n", 95, 95&(95+1))
	fmt.Printf("把右起第一个0变成1 | %b -> %b \n", 159, 159|(159+1))
	fmt.Printf("把右边连续的0变成1 | %b -> %b \n", 88, 88|(88-1))
	fmt.Printf("取右边连续的1 | %b -> %b \n", 159, 159^(159+1))
	fmt.Printf("去掉右起第一个1的左边 | %b -> %b \n", 88, 88&^(88-1))
	reads(1314520)
}

//位运算方式 交换整数
func swap(a, b *int) {
	if *a == *b {
		return
	}
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

//基本运算交换整数
func swap2(a, b int) {
	a = a + b
	b = a - b
	a = a - b
	fmt.Print(a, b, "\n")
}

// 互换两个变量的值
// 不需要使用第三个变量做中间变量
func swap3(a, b int) (int, int) {
	a ^= b // 异或等于运算
	b ^= a
	a ^= b
	return a, b
}

//统计一个整数n的二进制表示中1的个数与0的个数
func countOne(n int) (i, o int) {
	i = 0
	o = 0
	for n > 0 {
		if (n & 1) == 1 {
			i++
		} else {
			o++
		}
		n = n >> 1
	}
	return
}

//用位运算实现求绝对值－有效避开if-else判断
func abs_bit(n int32) (o int32) {
	o = n >> 31
	return (n + o) ^ o
}

// 获取0-n之间的所有偶数
func even(a int) (array []int) {
	for i := 0; i < a; i++ {
		if i&1 == 0 { // 位操作符&与C语言中使用方式一样
			array = append(array, i)
		}
	}
	return array
}

// 左移、右移运算
func shifting(a int) int {
	a = a << 1
	a = a >> 1
	return a
}

// 变换符号
func nagation(a int) int {
	// 注意: C语言中是 ~a+1这种方式
	return ^a + 1 // Go语言取反方式和C语言不同，Go语言不支持~符号。
}
