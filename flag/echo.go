package main

import (
	"flag"
	"os"
)

var newline = flag.Bool("n", false, "print on newline")
var appendstr = flag.Bool("o", false, "print append")

const (
	Space     = " "
	LineBreak = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	var s string = "iIII"
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if *newline {
		s += LineBreak
	}
	if *appendstr {
		s += "ooooo"
	}
	os.Stdout.WriteString(s)
}
