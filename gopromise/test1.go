package main

import (
	"fmt"
	promise "github.com/fanliao/go-promise"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	task := func() (r interface{}, err error) {
		url := "http://www.baidu.com"

		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	}

	f := promise.Start(task).OnComplete(func(v interface{}) {
		fmt.Println("ok...\n")
	}).OnSuccess(func(v interface{}) {
		fmt.Println("success")
	})
	f.Get()

	task1 := func() (r interface{}, err error) {
		return "ok1", nil
	}

	task2 := func() (r interface{}, err error) {
		return "ok2", nil
	}

	f = promise.WhenAll(task1, task2)
	v, _ := f.Get()
	fmt.Printf("%s", v)
	i := 0
	task3 := func(canceller promise.Canceller) (r interface{}, err error) {
		for i < 50 {
			if canceller.IsCancellationRequested() {
				canceller.Cancel()
				return 0, nil
			}
			time.Sleep(100 * time.Millisecond)
		}
		return 1, nil
	}
	f = promise.Start(task3)
	f.RequestCancel()

	f.Get()                      //return nil, promise.CANCELLED
	fmt.Println(f.IsCancelled()) //true

}
