package main

import (
	//"bytes"
	"fmt"
	//"github.com/fzzy/radix/extra/pubsub" // 发布订阅不太好用
	//"github.com/fzzy/radix/redis" //发布订阅在 extra包中
	//"github.com/hoisie/redis" //还有一些命令未实现,如watch, unwatch
	redis "github.com/alphazero/Go-Redis" //这个比较好用
	"os"
	"time"
)

func main() {
	// 连接redis
	var spec *redis.ConnectionSpec = new(redis.ConnectionSpec)
	spec.Host("127.0.0.1")
	spec.Port(6379)
	spec.Heartbeat(time.Duration(5) * time.Second)

	redisClient, err := redis.NewAsynchClientWithSpec(spec)
	checkErr(err)
	client, err := redis.NewPubSubClientWithSpec(spec)
	checkErr(err)
	client.Subscribe("htime.me")

	redisClient.Publish("htime.me", []byte("哈哈哈"))

	chans := client.Messages("htime.me")
	for {
		select {
		case s := <-chans:
			fmt.Println(string(s))
		}
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
