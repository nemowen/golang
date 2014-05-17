package main

import (
	"gotest/greentea/serial"
	"io"
	"log"
	"runtime"
	"time"
)

var (
	com    io.ReadWriteCloser
	buffer = make([]byte, 0, 128)
	ok     chan int
)

func init() {
	c := &serial.Config{Name: "COM2", Baud: 128000}
	ok = make(chan int, 10)
	var err error
	com, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	exit := make(chan bool)
	go read()
	go parse()

	<-exit
}

func parse() {
	for {
		select {
		case <-ok:
			log.Printf("%p %d %d\n", buffer, len(buffer), cap(buffer))

		case <-time.After(3 * time.Second):
			buffer = append(buffer[:0])
			log.Println("清空缓存 ")
			log.Printf("%p %d %d\n", buffer, len(buffer), cap(buffer))

		}
	}

}

func read() {
	var inbyte = make([]byte, 10)
	for {
		n, err := com.Read(inbyte)
		if err != nil {
			log.Fatal(err)
		}
		buffer = append(buffer, inbyte[0:n]...)
		if n > 1 {
			ok <- n
		}
	}
}
