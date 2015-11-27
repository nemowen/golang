package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

const lenPath = len("/view/")

var (
	titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")
	templaters     = make(map[string]*template.Template)
)

func init() {
	for _, name := range []string{"view", "edit"} {
		templaters[name] = template.Must(template.ParseFiles(name + ".html"))
	}
}

func main() {
	http.HandleFunc("/view/", makeHandle(viewHandle))
	http.HandleFunc("/edit/", makeHandle(editHandle))
	http.HandleFunc("/save/", makeHandle(saveHandle))
	http.ListenAndServe(":8080", nil)
}

func makeHandle(fu func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[lenPath:]

		if !titleValidator.MatchString(title) {
			http.Error(w, "title is invalid", http.StatusForbidden)
			return
		}
		fu(w, r, title)
	}
}

func viewHandle(w http.ResponseWriter, r *http.Request, title string) {
	p, err := load(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func editHandle(w http.ResponseWriter, r *http.Request, title string) {
	p, err := load(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandle(w http.ResponseWriter, r *http.Request, title string) {
	r.ParseForm()
	content := r.Form.Get("Body")
	if len(content) < 1 {
		http.Error(w, "please input something...", http.StatusForbidden)
	}
	err := ioutil.WriteFile(title+".txt", []byte(content), os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templaters[tmpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
