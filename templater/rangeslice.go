package main

import (
	"html/template"
	"net/http"
	"time"
)

type Book struct {
	Author  string
	Content string
	Date    time.Time
}

var htmlstr = `
	{{range .}}
		{{with .Author}}
			<p><b>{{html .}}</b> wrote</p>
		{{else}}
			<p>An anonymous person wrote</p>	
		{{end}}
		
		<pre>{{html .Content}}</pre>
		<pre>{{html .Date}}</pre>
	{{end}}
`

var bookList = []*Book{
	&Book{"Nemo", "the go programming", time.Now()},
	&Book{"nora", "how to study english", time.Now()},
}

func main() {
	t := template.Must(template.New("html").Parse(htmlstr))
	http.HandleFunc("/book/", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, bookList)
	})
	http.ListenAndServe(":8080", nil)
}
