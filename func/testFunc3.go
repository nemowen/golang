package main

func sum(name string, args ...int) {
	var n int

	for _, s := range args {
		n += s
	}
	println(name, n)
}

func sumslice(name string, args []int) {
	var n int
	for _, s := range args {
		n += s
	}
	println(name, n)
}

// 闭包支持
func closures(x int) func(int) int {
	// 返回匿名函数
	return func(y int) int {
		return x + y
	}
}

// 闭包指向同⼀一个变量，而不是复制。
func test(x int) func() {
	println(&x, x)
	return func() {
		println(&x, x)
	}
}

func main() {
	sum("test1:", 1, 2, 3, 4, 5, 5)
	sumslice("test2:", []int{1, 2, 3, 4, 5, 5})

	f := closures(20)
	println(f(10))
	println(f(20))

	var fs []func(int) int
	for i, j := 0, 3; i < j; i++ {
		fs = append(fs, func(x int) int {
			return x + 10
		})
	}

	for i, fu := range fs {
		println(fu(i))
	}

	test(100)
}
