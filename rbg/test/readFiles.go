package main

import (
	"log"
	"os"
)

func main() {
	f, e := os.Create("D:/EN/flag.ini")
	if e != nil {
		log.Println("打开文件失败：", "D:/EN/flag.ini")
		return
	}
	defer f.Close()
	f.WriteString("OK")

}
