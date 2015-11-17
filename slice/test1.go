package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var urls = []string{
		"http://www.baidu.com",
		"http://www.sina.com.cn",
		"http://www.qq.com",
	}

	exit := make(chan bool)
	processed := make(chan string, 5)

	for _, url := range urls {
		go func(url string) {
			rep, err := http.Get(url)
			checkerr(err)
			b, err := ioutil.ReadAll(rep.Body)
			checkerr(err)
			processed <- string(b)
			rep.Body.Close()
		}(url)
	}

	var sum int = 0
	for {
		select {
		case str := <-processed:
			fmt.Println(str)
			sum++
			if sum >= 3 {
				exit <- true
			}
		}
	}

	<-exit
	fmt.Println(">>>>>>>>>>>>>Done", sum)

}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
