package main

type User struct {
	id   int
	name string
}

func test(i interface{}) {
	switch v := i.(type) {
	case *User:
		println("User: Name = ", v.name)
	case string:
		println("string:", v)
	case int:
		println("int:", v)
	default:
		println("default:", v)
	}
}

func main() {
	user := &User{1, "Nora"}
	test(user)
	str := "it's a string"
	test(str)
	in := "it's a int"
	test(in)
	test(float64(12))
}
