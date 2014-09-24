package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func (this *User) test(name string) string {
	fmt.Println("I'M " + name)
	return name
}

func (this User) test2(name string) string {
	fmt.Println("2 I'M " + name)
	return name
}

func main() {
	u := User{1, "Nemo"}
	u.test("name")
}
