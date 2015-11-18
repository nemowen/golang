package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type A struct {
	X, Y, Z int32
	Name    string
}

type B struct {
	X, Z, y int64
	Name    string
}

func main() {
	var connect bytes.Buffer
	var err error

	enc := gob.NewEncoder(&connect)
	dec := gob.NewDecoder(&connect)

	err = enc.Encode(A{12, 13, 14, "NEMO"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	var b B
	err = dec.Decode(&b)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	fmt.Printf("%q:{X:%d, y:%d, Z:%d}\n", b.Name, b.X, b.y, b.Z)

}
