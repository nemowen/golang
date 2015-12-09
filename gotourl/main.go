package main

import (
	"flag"
	"fmt"
	"github.com/nemowen/golang/gotourl/store"
	"html/template"
	"net/http"
	"net/rpc"
)

var (
	urlstore  *URLStore
	urlTmpl   *template.Template
	rpcEnable = flag.Bool("rpc", "false", "enable RPC Server")
)

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
	flag.Parse()
	urlstore = store.NewURLStore("store.json")
	if *rpcEnable {
		rpc.RegisterName("store", urlstore)
		rpc.HandleHTTP()
	}
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8088", nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	var url string
	err := urlstore.Get(&key, &url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	var key string
	err := urlstore.Put(&url, &key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "http://%s/%s", r.Host, key)
}
