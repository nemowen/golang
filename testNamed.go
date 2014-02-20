package main

type MyInt int
type MyIntSlice []int

func main() {
	var a MyInt = 10
	// var b int = a //cannot use a (type MyInt) as type int in assignment
	var b int = int(a)
	// var c MyInt = b //cannot use b (type int) as type MyInt in assignment
	var c MyInt = MyInt(b)
	println(a, b, c)

	// 而当目标是未命名类型时，无需显式转换。

	var d MyIntSlice = []int{1, 2, 3}
	var e []int = d
	println(d, e)

	f := MyIntSlice{
		1,
		2,
	}
	println(f)
}
