package main

import (
	"fmt"
)

type Request struct {
	a, b  int
	reply chan int
}

type opFunc func(a, b int) int

func run(op opFunc, req *Request) {
	req.reply <- op(req.a, req.b)
}

func server(op opFunc, service chan *Request) {
	for {
		req := <-service
		go run(op, req)
	}
}

func startServer(op opFunc) chan *Request {
	req := make(chan *Request)
	go server(op, req)
	return req
}

func main() {
	reqChan := startServer(func(a, b int) int { return a + b })
	const N = 100
	//var reqs [N]Request
	reqs := make([]*Request, 0, 100)
	for i := 0; i < N; i++ {
		req := new(Request)
		req.a = i
		req.b = i + N
		req.reply = make(chan int)
		reqs = append(reqs, req)
		reqChan <- req
	}

	for i, v := range reqs {
		if c := <-v.reply; c != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, "is ok!", c)
		}
		fmt.Println("done")
	}

	//	for i := N - 1; i >= 0; i-- {
	//		range reqs
	//		if c := <-reqs[i].reply; c != N+2*i {
	//			fmt.Println("fail at", i)
	//		} else {
	//			fmt.Println("Request ", i, "is ok!", c)
	//		}
	//		fmt.Println("done")
	//	}
}
