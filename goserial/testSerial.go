package main

import (
	//"bytes"
	"github.com/tarm/goserial"
	"io"
	"log"
	"runtime"
)

var (
	com    io.ReadWriteCloser
	buffer = make([]byte, 128)
)

func init() {
	c := &serial.Config{Name: "COM1", Baud: 128000}
	var err error
	com, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	exit := make(chan bool)

	<-exit
}

func read() {
	n, err := com.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n, string(buffer[0:n]))
}
