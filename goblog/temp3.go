package main

import (
	"io"
	"net/http"
	"time"
)

var router map[string]func(w http.ResponseWriter, r *http.Request)

func main() {
	router = make(map[string]func(w http.ResponseWriter, r *http.Request))
	router["/hello"] = sayHello
	router["/index"] = index
	server := http.Server{
		Addr:         ":8080",
		Handler:      &MyHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if f, ok := router[r.URL.String()]; ok {
		f(w, r)
		return
	}

	io.WriteString(w, r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My dear, I miss you!")
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My dear, I miss you! go home")
}
