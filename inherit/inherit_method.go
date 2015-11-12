package main

import (
	"fmt"
)

type Base struct {
	id int64
}

func (b *Base) Id() int64 {
	return b.id
}

func (b *Base) SetId(id int64) {
	b.id = id
}

type Person struct {
	FirstName string
	LastName  string
	Base
}

type Employee struct {
	salary float32
	Person
}

func main() {
	// 如果要写字段名就全写字段名，Person : Person{},如果不写，就全不写，如下
	e := &Employee{23.6, Person{"Wen", "Nemo", Base{1}}}
	// or
	// e := &Employee{salary: 23.6, Person: Person{FirstName: "Nemo", LastName: "Wen", Base: Base{id: 1}}}

	e.SetId(5)
	fmt.Println(e.Id())
}
