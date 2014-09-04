/*****************************************************
	利用反射显示一个 "未知" 结构类型的字段、方法等信息
*****************************************************/
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
	Job  string
}

func (this User) SayHi(msg string) {
	fmt.Printf("Hi My name is ", this.Name)
	fmt.Printf("Nice to meet you!!!")
}

func showObject(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	// Type kind
	if k := t.Kind(); k != reflect.Struct {
		fmt.Printf("%v of type kind is not a struct!\n", k)
		return
	}

	// Type Name
	fmt.Printf("Type:%s\n", t.Name())

	// Fields
	fmt.Printf("Fields:\n")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("	%6s: %v = %v\n", field.Name, field.Type, value)
	}

	//Methods
	fmt.Printf("Methods:\n")
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("	%6s: %v\n", method.Name, method.Type)
	}

}

func main() {
	u := User{1, "Nemo", 27, "Programmer"}
	showObject(u)
	ss := "ok"
	showObject(ss)
}
