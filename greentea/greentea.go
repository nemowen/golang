package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gotest/greentea/serial"
	"gotest/greentea/tools"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type configObj struct {
	ComName       string        // 串口名称
	Baud          int           // 波特率
	IniSavePath   string        // ini文件保存路径
	BmpSavePath   string        // bmp保存路径
	BmpDaysToKeep time.Duration // bmp保存天数
}

const (
	jsons         string = "D:/PROGRAM/GO/Development/src/gotest/greentea/config.json" // 客户端配置文件路径
	DATA_END_FLAG string = "*s[OCR_End]s*"                                             // 本笔结束

	I_S_DATE_FLAG, I_E_DATE_FLAG string = "*d{", "}d*"   //数据日期标识位
	I_S_TIME_FLAG, I_E_TIME_FLAG string = "*t{", "}t*"   //数据时间标识位
	I_S_NO_FLAG, I_E_NO_FLAG     string = "*no{", "}no*" //数据顺序号标识位
	I_S_BN_FLAG, I_E_BN_FLAG     string = "*bn{", "}bn*" //数据冠字号标识位

	M_S_DATE_FLAG, M_E_DATE_FLAG string = "*d{", "}d*" //机器状态：数据日期标识位
	M_S_TIME_FLAG, M_E_TIME_FLAG string = "*t{", "}t*" //机器状态：数据时间标识位

	STATUS_INIT      int = 1 // 初始化工作
	STATUS_READ_DONE int = 2 // 读取完成
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

	// 检查并设置默认值
	if config.IniSavePath == nil {
		config.IniSavePath = "C:/CNRData"
	}
	if config.BmpSavePath == nil {
		config.BmpSavePath = "D:/SNRData"
	}

	c := &serial.Config{Name: config.ComName, Baud: config.Baud}
	var err error
	com, err = serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ok = make(chan int, 50)
	exit := make(chan bool)
	bmpClearTask := tools.NewTask("bmpClearTask", "50 55 22 * * * ", bmpClear)
	tools.AddTask("bmpClearTask", bmpClearTask)
	tools.StartTask()
	defer tools.StopTask()

	go read()
	go parse()

	<-exit
}

func read() {
	var inbyte = make([]byte, 1024)
	for {
		n, err := com.Read(inbyte)
		if err != nil {
			fmt.Printf("%s", err)
		}
		buffer = append(buffer, inbyte[0:n]...)
		if bytes.Contains(buffer, "*s[start]s*") {
			ok <- STATUS_INIT
		}
		if bytes.Contains(buffer, []byte(DATA_END_FLAG)) {
			ok <- STATUS_READ_DONE
		}
	}
}

func parse() {
	for {
		select {
		case status := <-ok:
			//fmt.Printf("%s %d %d\n", buffer, len(buffer), cap(buffer))
			if STATUS_INIT == status {
				now := time.Now().Format("20060102")
				if now != currentDay { // 每日清空交易笔数
					countTimesDay = 0
					currentDay = now
				}

			} else if STATUS_READ_DONE == status {
				inifile := os.OpenFile(config.IniSavePath, os.O_WRONLY|os.O_CREATE, 0666)

				// 解析开始

			}

		case <-time.After(5 * time.Second):
			countTimesDay += 1 // 统计交易笔数
			ctd := strconv.Itoa(countTimesDay)
			path := filepath.Join(config.BmpSavePath, currentDay, ctd)
			os.MkdirAll(path, 0666)

			for i := 1; i <= 10; i++ { // 测试写入BMP
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

func bmpClear() error {
	// 定时检查过期数据
	files, err := ioutil.ReadDir(config.BmpSavePath)
	if err != nil {
		return errors.New("未找到BMP目录：" + err)
	}
	for _, file := range files {
		filename := file.Name()
		t, e := time.Parse("20060102", filename)
		if e != nil {
			fmt.Println("目录名称不正确", e)
			continue
		}
		if time.Now().Sub(t.Add(config.BmpDaysToKeep*24*time.Hour)) > 0 {
			os.RemoveAll(filepath.Join(config.BmpSavePath, filename))
		}
	}
	return nil
}
