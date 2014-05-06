package main

import (
	"gotest/rbg/server/rpcobj"
	"gotest/rbg/server/utils"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"runtime"
)

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	exit := make(chan bool)

	u := new(rpcobj.Obj)
	rpc.Register(u)

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", utils.Server_preferences.SERVER_IP_PORT)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()
	go http.Serve(l, nil)

	log.Println("服务器已经启动！")
	<-exit
}
