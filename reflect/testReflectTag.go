package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `field:"username" type:"nvarchar(20)"`
	Age  int    `field:"age" type:"tinyint"`
	Tel  string
}

func main() {
	var u User
	t := reflect.TypeOf(u)
	f, _ := t.FieldByName("Name")
	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("field"))
	fmt.Println(f.Tag.Get("type"))
}
