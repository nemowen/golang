package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
}

func main() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{.Name}}!\n")
	p := Person{Name: "nemo"}
	if e := t.Execute(os.Stdout, p); e != nil {
		fmt.Println("There was an error:", e)
	}
}
