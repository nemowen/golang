package main

import (
	"html/template"
	"os"
)

var (
	mone = `this is template one: 11111 {{template "all"}}`
	mtwo = `this is template two: 22222 {{template "m1"}}`
	mall = `show all template:
			1>>>>: {{template "m1"}}
			2>>>>: {{template "m2"}}
			`
)

func main() {
	tmpl := template.Must(template.New("m1").Parse(mone))
	tmpl = template.Must(tmpl.New("m2").Parse(mtwo))
	tmpl = template.Must(template.New("all").Parse(mall))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
