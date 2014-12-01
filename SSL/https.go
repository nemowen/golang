package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, err := client.Get("https://kyfw.12306.cn/otn/login/init")
	if err != nil {
		fmt.Println(err)
	}
	if r.StatusCode != 200 {
		fmt.Println("connect to 12306 error")
	}
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
}
