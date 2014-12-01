package main

import (
	"fmt"
)

// Do not communicate by sharing memory. Instead, share memory by communicating。

type Data struct {
	args []int
	ch   chan string
}

func Server() chan *Data {
	// 创建一个信道
	reqs := make(chan *Data)
	go func() {
		// 迭代取出信道里的值
		for data := range reqs {
			// 取出的值交给服务器处理
			go serverProess(data)
		}
	}()

	return reqs
}

func serverProess(data *Data) {
	s
	x := 0
	for v := range data.args {
		x += v
	}
	s := fmt.Sprintf("server:%d", x)
	data.ch <- s
}

func main() {
	server := Server()
	data := &Data{[]int{1, 2, 3, 4, 5}, make(chan string)}
	server <- data
	fmt.Println(<-data.ch)
	close(server)
}
