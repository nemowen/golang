package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func cancelHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	time.Sleep(10 * time.Second)
	cancel()
}

func timeoutHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//go doStuff(ctx)
	go doTimeoutStuff(ctx)

	time.Sleep(10 * time.Second)
	cancel()
}

func doTimeoutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		if deadline, ok := ctx.Deadline(); ok {
			logg.Println("deadline set")
			if time.Now().After(deadline) {
				logg.Println("-->", ctx.Err().Error())
				return
			}
		}

		select {
		case <-ctx.Done():
			logg.Println("done")
			return
		default:
			logg.Println("work")
		}
	}
}

func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Println("done")
			return
		default:
			logg.Println("work")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	timeoutHandler()
	logg.Println("down")
}
