package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) SayHi() {
	fmt.Println("hello ", u.Name)
}

func main() {
	u := User{"nemo", 27}
	fmt.Println(reflect.TypeOf(u).MethodByName("SayHi"))

}

// func info(o interface{}) {
// 	t := reflect.TypeOf(o)
// 	v := reflect.ValueOf(o)

// 	for i := 0; i < t.NumField(); i++ {

// 	}

// }
