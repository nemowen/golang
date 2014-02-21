package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	UserName string
}

type SingleHost struct {
	handler     http.Handler
	allowedHost string
}

func (s *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	fmt.Println(host + ":" + s.allowedHost)
	if host == s.allowedHost {
		s.handler.ServeHTTP(w, r)
	} else {
		fmt.Fprintln(w, "Error")
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	n := r.Form["username"][0]
	if "nemo" == n {
		fmt.Println(n, "is login.")
		fmt.Fprintln(w, "Hi "+n+"welcome to here...!")
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func NewSingleHost(handler http.Handler, allowedHost string) *SingleHost {
	return &SingleHost{handler: handler, allowedHost: allowedHost}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                    //解析参数，默认是不会解析的
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/login.html")
		p := Person{UserName: "Nemo"}
		t.Execute(w, p)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	//http.HandleFunc("/", sayhelloName) //设置访问的路由
	single := NewSingleHost(http.HandlerFunc(login), "127.0.0.1")
	err := http.ListenAndServe(":80", single) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
