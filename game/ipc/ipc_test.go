package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(request string) string {
	return "ECHO:" + request
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIPC(t *testing.T) {
	server := NewIPCServer(&EchoServer{})
	client1 := NewIPCClient(server)
	client2 := NewIPCClient(server)

	resp1 := client1.Call("From Client1")
	resp2 := client2.Call("From Client2")

	if resp1 != "ECHO:From Client1" || resp2 != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}
	client1.Close()
	client2.Close()
}
