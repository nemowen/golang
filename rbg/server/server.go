package main

import (
	"gotest/rbg/server/rpcobj"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"runtime"
)

func main() {
	exit := make(chan bool)
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	u := new(rpcobj.Obj)
	rpc.Register(u)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1314")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()
	go http.Serve(l, nil)

	log.Println("服务器已经启动！")
	<-exit
}
