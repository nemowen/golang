package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	f, e := os.Create("D:/EN/flag.ini")
	if e != nil {
		log.Println("打开文件失败：", "D:/EN/flag.ini")
		return
	}
	defer f.Close()
	f.WriteString("OK")

	dayLastYear, _ := strconv.Atoi(time.Now().Add(time.Hour * -(1 * 24)).Format("20060102"))
	fmt.Printf("%v", dayLastYear)

}
