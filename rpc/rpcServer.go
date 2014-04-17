package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	. "rpctest"
	"time"
)

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8123")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()
	go http.Serve(l, nil)
	time.Sleep(60 * time.Second)
}
