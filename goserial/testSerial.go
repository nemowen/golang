package main

import (
	"bytes"
	"github.com/tarm/goserial"
	"io"
	"log"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	c := &serial.Config{Name: "COM2", Baud: 128000}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	exit := make(chan bool, 1)
	ioch := make(chan []byte, 128)

	go readData(&ioch, s)

	go func(ioch *chan []byte) {
		for {
			a := <-*ioch
			log.Println(string(a))
		}

		//exit <- true
	}(&ioch)

	<-exit

}

func readData(ioch *chan []byte, s io.ReadWriteCloser) {
	for {
		data := make([]byte, 1)
		_, err := s.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		*ioch <- data
	}
}
