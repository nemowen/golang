package main

import (

	//"io"
	"log"
	"os"
)

func main() {

	b := make([]byte, 1024, 1024)
	log.Println("init:", len(b))
	f, _ := os.Open("d:/7.bmp")
	f.Read(b)
	log.Print(b, "now:", cap(b))
}
