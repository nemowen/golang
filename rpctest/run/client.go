package main

import (
	"fmt"
	"gotest/rpctest"
	"image"
	"log"
	"net/rpc"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	client, e := rpc.DialHTTP("tcp", "192.168.0.109:1314")
	if e != nil {
		log.Fatal("dialing:", e)
	}
	u := new(rpctest.User)
	var a int = 1
	exit := make(chan bool)
	go func() {
		for i := 10000; i > 0; i-- {
			// go method: 异步方法 , 似乎效率没有call同步方法高
			// ch := client.Go("User.GetUserById", &a, u, nil)
			// reply := <-ch.Done
			// if u, ok := reply.Reply.(*rpctest.User); ok {
			// 	fmt.Println(u.Name, u.Age, u.Job)
			// } else {
			// 	break
			// 	exit <- true
			// }

			// call method: 同步方式，似乎效率比go 异步方法高
			err := client.Call("User.GetUserById", &a, u)
			if err == nil {
				fmt.Println(u.Name, u.Age, u.Job)
			} else {
				exit <- true
			}
			log.Println(">>>>>>", i)
		}
		exit <- true
	}()
	<-exit

}
