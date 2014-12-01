package main

// 多个返回值
func swap(a, b int) (int, int) {
	return b, a
}

// 命名返回参数
func change(a, b int) (x, y int) {
	x = a + 100
	y = b + 100
	return
}

func test(a int) (x int) {
	if x := 10; a > 0 {
		return x + 1
	}

	return
}

func main() {
	b, a := swap(1, 2)
	println(b, a)
	y, x := change(3, 4)
	println(y, x)
	r := test(10)
	println(r)
}
