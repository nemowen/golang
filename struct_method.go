package main

import "fmt"

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

func (p *Person) SayHi() {
	fmt.Printf("Hi, I am %s, you can call me on %s\n", p.name, p.phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	//example 1:
	s1 := Student{Person{"Nora", 26, "16548756877"}, "天河大学"}
	//example 2:
	e2 := Employee{Person{"Nemo", 26, "18665069119"}, "软件工程师", "天创时尚"}

	s1.SayHi()
	e2.SayHi()
}
