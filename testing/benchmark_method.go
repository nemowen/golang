package main

import (
	"fmt"
	"testing"
)

type Person struct {
	name     string
	age      int
	position string
	addr     string
	phone    string
	sex      string
}

var person = Person{"Nemo", 28, "IT 软件工程师", "广州广东省", "18665069119", "男"}

func (p *Person) ShowInfo() {

}

func (p Person) Info() {

}

func BenchmarkPersonPoint(t *testing.B) {
	for i := 0; i < t.N; i++ {
		person.ShowInfo()
	}
}

func BenchmarkPersonValue(t *testing.B) {
	for i := 0; i < t.N; i++ {
		person.Info()
	}
}

func main() {
	fmt.Println("Person-Point:", testing.Benchmark(BenchmarkPersonPoint).String())
	fmt.Println("Person-value:", testing.Benchmark(BenchmarkPersonValue).String())
}
