package main

import (
	"html/template"
	"os"
)

func main() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $3:=`ok`}}{{$3}}\n{{end}}"))
	t.Execute(os.Stdout, nil)
	t = template.New("test1")
	t = template.Must(t.Parse("{{with $4:=`123`}}{{$4}}\n{{end}}"))
	t.Execute(os.Stdout, nil)
}
