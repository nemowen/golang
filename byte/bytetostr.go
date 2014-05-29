package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"
	"os"
	//"unicode/utf16"
	//"unicode/"
	//"unicode/utf8"
)

func main() {
	f, _ := os.OpenFile("e:/archive_.txt", os.O_CREATE|os.O_APPEND, 0666)
	b, e := ioutil.ReadAll(f)
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Printf("%s", string(b))
	f.Close()
}
