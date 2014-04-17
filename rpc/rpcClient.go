package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpctest"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8123")
	if err != nil {
		log.Fatal("arith error:", err)
	}
	args := &rpctest.Args{7, 8}
	var reply int

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
