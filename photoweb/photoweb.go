package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

type User struct {
	Name string
}

const (
	UPLOAD_DIR = "./upload"
)

func main() {
	cpus := runtime.NumCPU()
	maxps := runtime.GOMAXPROCS(cpus)
	fmt.Println(cpus, ":", maxps)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", myHandle)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func myHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/upload.html")
		p := User{Name: "Nemo"}
		t.Execute(w, p)
		return
	}
	if r.Method == "POST" {
		// 获取文件流
		file, fh, err := r.FormFile("image")
		parseErr(w, err)

		// 获取文件名称
		filename := fh.Filename
		// 关闭文件流
		defer file.Close()

		//创建一个文件
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		parseErr(w, err)

		n, err := io.Copy(t, file)
		log.Print("size:", n)

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func parseErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
