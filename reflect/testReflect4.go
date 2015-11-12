package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age      int
}
type Admin struct {
	User
	title string
}

func main() {
	u := Admin{User{"Jack", 23}, "NT"}
	v := reflect.ValueOf(u)
	fmt.Println(v.FieldByName("title").String())
	fmt.Println(v.FieldByName("age").Int())
	fmt.Println(v.FieldByIndex([]int{0, 1}).Int())

	agev := v.FieldByName("age")
	if agev.CanInterface() {
		fmt.Println(agev.Interface())
	} else {
		fmt.Println(agev.Int())
	}

}
