package main

import (
	"html/template"
	"os"
)

func main() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $x:=`hello`}}{{printf `----%s ----%s` $x `Mary`}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
}
