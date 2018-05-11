package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,el

const (
	httpPort = "8090"
)

func PrintMessage(w http.ResponseWriter, r *http.Request) {
	p := message.NewPrinter(language.English)
	p.Fprintf(w, "Hello, %v", html.EscapeString(r.Host))
}

func main() {
	var port string
	flag.StringVar(&port, "port", httpPort, "http port")
	flag.Parse()

	server := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 16,
		Handler:        http.HandlerFunc(PrintMessage)}

	log.Fatal(server.ListenAndServe())
}
