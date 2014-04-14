package main

import (
	"gotest/rpctest"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func inServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的

	//fmt.Fprintf(w, " 222 ") //这个写入到w的是输出到客户端的
}

func main() {
	u := new(rpctest.User)
	rpc.Register(u)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1314")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()

	go http.Serve(l, nil)

	time.Sleep(10 * time.Minute)

}
