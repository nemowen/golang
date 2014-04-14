package main

import (
	"fmt"
	"gotest/rpctest"
	"log"
	"net/rpc"
)

func main() {
	client, e := rpc.DialHTTP("tcp", "127.0.0.1:1314")
	if e != nil {
		log.Fatal("dialing:", e)
	}
	u := new(rpctest.User)
	var a int = 2
	ch := client.Go("User.GetUserById", &a, u, nil)
	reply := <-ch.Done
	if u, ok := reply.Reply.(*rpctest.User); ok {
		fmt.Println(u.Name, u.Age, u.Job)
	}

}
