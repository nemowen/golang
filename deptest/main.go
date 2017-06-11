package main

import (
	"github.com/beego/mux"
	"log"
	"net/http"
)

func main() {
	mx := mux.New()
	mx.Handler("GET", "/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", mx))

	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	//if err := http.ListenAndServe(":1987", nil); err != nil {
	//	log.Fatal("ListenAndServe:", err)
	//}
}
