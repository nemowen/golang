package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IPCServer struct {
	Server
}

func NewIPCServer(server Server) *IPCServer {
	return &IPCServer{server}
}

func (server *IPCServer) Connect() chan string {
	session := make(chan string)

	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b) // 返回结果
		}
		fmt.Println("Session closed.")
	}(session)
	fmt.Println("A new session has been created successfully.")
	return session
}
