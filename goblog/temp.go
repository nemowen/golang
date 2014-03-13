package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello Nemo, this is version 1")
}
