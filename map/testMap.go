package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
	age  int
}

func main() {
	us := make(map[string]User)
	u := User{id: 1, name: "Nemo", age: 27}
	us["wo"] = u

	su := make(map[User]string)
	su[u] = "Nora"

	fmt.Println(us, us["wo"].age)
	fmt.Println(su, su[u])
}
