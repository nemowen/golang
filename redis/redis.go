package main

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"os"
	"time"
)

func main() {
	// 连接redis
	client, err := redis.DialTimeout("tcp", "192.168.202.128:6379", time.Duration(10)*time.Second)
	checkErr(err)
	defer client.Close()

	// 验证
	r := client.Cmd("auth", "wenbin")

	// 选择数据库
	r = client.Cmd("select", 0)

	// 清空当前数据库
	//r = client.Cmd("flushdb")

	// 写入一个Key:aa value:aa
	r = client.Cmd("set", "aa", "aa")

	// 取得key:aa 的值
	s, _ := client.Cmd("get", "aa").Str()
	fmt.Println("echo:", s)

	r = client.Cmd("hgetall", "myhash")
	for id, v := range r.Elems {
		fmt.Printf("id:%d, %s\n", id, v.String())
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
