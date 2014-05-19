package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotest/greentea/serial"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type configObj struct {
	ComName       string // 串口名称
	Baud          int    // 波特率
	IniSavePath   string // ini文件保存路径
	BmpSavePath   string // bmp保存路径
	BmpDaysToKeep int    // bmp保存天数
}

const (
	jsons string = "D:/PROGRAM/GO/Development/src/gotest/greentea/config.json" // 客户端配置文件路径

	END_FLAG = "*s[OCR_End]s*" // 本笔结束
)

var (
	com           io.ReadWriteCloser
	buffer        = make([]byte, 0, 6<<10)
	ok            chan int
	config        *configObj
	countTimesDay int    // 当天交易次数
	currentDay    string // 今天日期
)

func init() {
	config = new(configObj)
	file, e := ioutil.ReadFile(jsons)
	if e != nil {
		fmt.Println("读取配置文件失败!请与管理员联系!")
		os.Exit(1)
	}
	json.Unmarshal(file, config)

	c := &serial.Config{Name: config.ComName, Baud: config.Baud}
	var err error
	com, err = serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ok = make(chan int, 50)
	exit := make(chan bool)

	go read()
	go parse()

	<-exit
}

func parse() {
	for {
		select {
		case <-ok:
			fmt.Printf("%s %d %d\n", buffer, len(buffer), cap(buffer))
		case <-time.After(5 * time.Second):
			countTimesDay += 1
			currentDay = time.Now().Format("20060102")
			ctd := strconv.Itoa(countTimesDay)
			path := filepath.Join(config.BmpSavePath, currentDay, ctd)
			os.MkdirAll(path, 0666)

			//os.Stat(name)
			//fmt.Println(path)
			for i := 1; i <= 10; i++ {
				files := filepath.Join(path, strconv.Itoa(i)+".bmp")
				fmt.Println(files)
				file, e := os.OpenFile(files, os.O_CREATE|os.O_WRONLY, 0666)
				if e != nil {
					fmt.Println(e)
				}
				file.Write(buffer)
				file.Close()
			}

		}
	}
}

func read() {
	var inbyte = make([]byte, 1024)
	for {
		n, err := com.Read(inbyte)
		if err != nil {
			fmt.Printf("%s", err)
		}
		buffer = append(buffer, inbyte[0:n]...)
		if bytes.Contains(buffer, []byte(END_FLAG)) {
			ok <- 1
		}
	}
}
