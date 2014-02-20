package main

import "fmt"
import "reflect"

type User struct {
	Id   int    `唯一标识`
	Name string `姓名`
	Age  int    `年龄`
}

func main() {
	u := User{Id: 1, Age: 27, Name: "Nemo"}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s (%s = %v)\n", f.Tag, f.Name, v.Field(i).Interface())

	}

}
