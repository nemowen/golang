package main

import (
	"fmt"
)

type Writer interface {
	write(s string)
}

type Reader interface {
	read() string
}

type WriterReader interface {
	Writer
	Reader
}

type WriterReaderTest struct {
}

func (*WriterReaderTest) write(s string) {
	fmt.Println("write:", s)
}

func (*WriterReaderTest) read() string {
	fmt.Println("read")
	return "Ok"
}

func main() {
	wr := &WriterReaderTest{}
	var ir Reader = wr
	var iw Writer = wr
	ir.read()
	iw.write("s232323")
}
