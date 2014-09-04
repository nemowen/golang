package main

/**
	使用Map,动态调用方法
**/
func main() {
	funcs := make(map[string]func() string, 10)
	funcs["a"] = func() string {
		println("this is func a")
		return "a"
	}
	funcs["b"] = func() string {
		println("this is func b")
		return "b"
	}
	println(funcs["a"]())
}
