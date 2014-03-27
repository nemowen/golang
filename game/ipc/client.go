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

func (client *IPCClient) Call(method, params string) (resp *Response, err error) {
	req := &Request(method, params)
	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}

	client.conn <- string(b)
	str := <-client.conn //等待返回值

	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1
}

func (client *IPCClient) Close() {
	client.conn <- "CLOSE"
}
