package main

import (
	"fmt"
	"github.com/nemowen/golang/gotourl/store"
	"html/template"
	"net/http"
)

var urlstore = store.NewURLStore("store.json")
var urlTmpl *template.Template

const UrlForm = `
	<html>
		<body>
			<h1>URLStore Size: {{.}}</h1>
			<form method="post" action="/add">
				url:<input name="url"/>
				<input type="submit" value="add"/>
			</form>
		</body>
	</html>
`

func init() {
	urlTmpl = template.Must(template.New("urlform").Parse(UrlForm))
}

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8088", nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := urlstore.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		urlTmpl.Execute(w, urlstore.Count())
		return
	}
	key := urlstore.Put(url)
	fmt.Fprintf(w, "http://%s/%s", r.Host, key)
}
