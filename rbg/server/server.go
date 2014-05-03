package main

import (
	"encoding/json"
	"gotest/rbg/config"
	"gotest/rbg/server/rpcobj"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"runtime"
)

var server_preferences config.ServerConfig

const (
	//客户端配置文件路径
	SERVER_PREFERENCES string = "C:/Windows/Server.Preferences.json"
)

func init() {
	//加载配置文件
	file, e := ioutil.ReadFile(SERVER_PREFERENCES)
	if e != nil {
		panic("读取配置文件失败！请与管理员联系！")
		os.Exit(1)
	}

	json.Unmarshal(file, &server_preferences)
}

func main() {
	exit := make(chan bool)
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	u := new(rpcobj.Obj)
	rpc.Register(u)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", server_preferences.SERVER_IP_PORT)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()
	go http.Serve(l, nil)

	log.Println("服务器已经启动！")
	<-exit
}
