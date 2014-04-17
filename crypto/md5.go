package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "test.txt"
	f, e := os.Open(file)
	defer f.Close()
	if e == nil {
		md5i := md5.New()
		io.Copy(md5i, f)
		fmt.Printf("%x %s\n", md5i.Sum([]byte("")), file)

		sha1i := sha1.New()
		io.Copy(sha1i, f)
		fmt.Printf("%x %s\n", sha1i.Sum([]byte("")), file)
	}

}
