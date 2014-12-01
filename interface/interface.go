package main

import (
	"fmt"
	"time"
)

type TInterface interface {
	ShowName()
	SayHi()
}

type Person struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Person
	school string
}

type Employee struct {
	Person
	jobname string
	company string
}

func (p Person) ShowName() {
	fmt.Printf("My name is %s\n", p.name)
}

func (s Student) SayHi() {
	fmt.Printf("Hi, I'm %s, I'm School name is %s\n", s.Person.name, s.school)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I'm %s, I'm Job is %s\n", e.Person.name, e.jobname)
}

func main() {
	a := time.Now().Second()
	student := Student{Person{"Nemo", 26, "13689564566"}, "暨南大学"}
	emp := Employee{Person{"Nora", 26, "13884564453"}, "销售员", "天创时尚"}

	var i TInterface

	i = student
	i.ShowName()
	i.SayHi()

	i = emp
	i.ShowName()
	i.SayHi()
	b := time.Now().Second()

	fmt.Println(b, "-", a)

}
