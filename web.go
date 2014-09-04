package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	UserName string
	IP       string
	Port     string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	username := r.FormValue("username")
	password := r.FormValue("password")

	if "nemo" == username && "123456" == password {
		fmt.Println(username, "is login.")
		fmt.Fprintln(w, "Hi "+username+" welcome to here...!")
	} else {

		t, _ := template.ParseFiles("./html/login.html")
		t.Execute(w, person)

	}

}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                    //解析参数，默认是不会解析的
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/login.html")
		t.Execute(w, person)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

var person = Person{UserName: "Nemo", IP: "127.0.0.1", Port: "8080"}

func main() {

	http.HandleFunc("/", login)                      //登录页面
	http.HandleFunc("/hello", sayhelloName)          //设置访问的路由
	err := http.ListenAndServe(":"+person.Port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
