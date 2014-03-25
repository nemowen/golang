package ipc

import (
	"encoding/json"
)

type IPCClient struct {
	conn chan string
}

func NewIPCClient(server *IPCServer) *IPCClient {
	c := server.Connect()
	return &IPCClient(c)
}
