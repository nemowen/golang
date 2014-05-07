package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	ts := "test123"
	md5Inst := md5.New()
	md5Inst.Write([]byte(ts))
	result := md5Inst.Sum([]byte(""))
	fmt.Printf("%X\n", result)

	shaInst := sha1.New()
	shaInst.Write([]byte(ts))
	result = shaInst.Sum([]byte(""))
	fmt.Printf("%X\n", result)
}
