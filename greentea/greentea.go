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
	LogSavePath   string        // log文件记录
	BmpDaysToKeep time.Duration // bmp保存天数
}

const (
	jsons string = "D:/PROGRAM/GO/Development/src/gotest/greentea/config.json" // 客户端配置文件路径

	DATA_S_FLAG, DATA_E_FLAG     = "*s[start]s*", "*s[output_end]s*" // 本笔结束
	I_S_DATE_FLAG, I_E_DATE_FLAG = "*d{", "}d*"                      // 数据日期标识位
	I_S_TIME_FLAG, I_E_TIME_FLAG = "*t{", "}t*"                      // 数据时间标识位
	I_S_NO_FLAG, I_E_NO_FLAG     = "*no{", "}no*"                    // 数据顺序号标识位
	I_S_BN_FLAG, I_E_BN_FLAG     = "*bn{", "}bn*"                    // 数据冠字号标识位
	M_S_DATE_FLAG, M_E_DATE_FLAG = "*d[", "]d*"                      // 机器状态：数据日期标识位
	M_S_TIME_FLAG, M_E_TIME_FLAG = "*t[", "]t*"                      // 机器状态：数据时间标识位

	STATUS_INIT      int = 1 // 初始化工作
	STATUS_READ_DONE int = 2 // 读取完成

	LineBreak = "\r\n" // windows 换行
)

var (
	com           io.ReadWriteCloser       // 串口对象
	buffer        = make([]byte, 0, 6<<10) // 缓冲区
	ok            chan int                 // 信号量
	config        *configObj               // 配置文件
	countTimesDay int                      // 当天交易次数
	currentDay    string                   // 今天日期
	bmpEndFlag    string                   // bmp数据结束标识
	snrinfo       *os.File                 // ini文件对象
	snrlog        *os.File                 // 日志记录
)

func init() {
	config = new(configObj)
	file, e := ioutil.ReadFile(jsons)
	if e != nil {
		fmt.Println("读取配置文件失败!请与管理员联系!")
		os.Exit(1)
	}
	json.Unmarshal(file, config)

	// 日志初始化
	log = logs.NewLogger(10000)
	// 日志文件记录
	log.SetLogger("file", `{"filename":"`+config.LogSavePath+`"}`)

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
	bmpClearTask := tools.NewTask("bmpClearTask", "59 59 23 * * * ", bmpClear)
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
		if bytes.Contains(buffer, []byte(DATA_S_FLAG)) {
			ok <- STATUS_INIT
		}
		if bytes.Contains(buffer, []byte(DATA_E_FLAG)) {
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

				// open or create SNRinfo.ini
				info, err := os.OpenFile(config.IniSavePath, os.O_CREATE|os.O_WRONLY, 0666)
				if err != nil {
					fmt.Println("创建SNRinfo.ini文件失败！", err)
				}
				snrinfo = info
				snrinfo.WriteString("[Cash_Info]" + LineBreak)
				snrinfo.WriteString("LEVEL4_COUNT=x" + LineBreak)
				snrinfo.WriteString("LEVEL3_COUNT=0" + LineBreak)
				snrinfo.WriteString("LEVEL2_COUNT=0" + LineBreak)
				snrinfo.WriteString("OperationTime=" + time.Now().Format("2006-01-02 15:04:05") + LineBreak)
				snrinfo.WriteString(LineBreak)

				now := time.Now().Format("20060102")
				if now != currentDay { // 每日清空交易笔数
					countTimesDay = 0
					currentDay = now
				}

				// open or create a log file
				snrlog, err = os.OpenFile(filepath.Join(config.LogSavePath, now+".log"),
					os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

			} else if STATUS_READ_DONE == status {
				//logfile := os.OpenFile(config.IniSavePath, os.O_WRONLY|os.O_CREATE, 0666)

				// create directory
				countTimesDay += 1 // 统计交易笔数
				ctd := strconv.Itoa(countTimesDay)
				path := filepath.Join(config.BmpSavePath, currentDay, ctd)
				os.MkdirAll(path, 0666)

				// parse machine state

				// parse info date
				i_date_s_index := bytes.Index(buffer, []byte(I_S_DATE_FLAG))
				i_date_e_index := bytes.Index(buffer, []byte(I_E_DATE_FLAG))
				i_date_data := buffer[i_date_s_index+len(I_S_DATE_FLAG) : i_date_e_index]

				// parse info time
				i_time_s_index := bytes.Index(buffer, []byte(I_S_TIME_FLAG))
				i_time_e_index := bytes.Index(buffer, []byte(I_E_TIME_FLAG))
				i_time_data := buffer[i_time_s_index+len(I_S_TIME_FLAG) : i_time_e_index]

				// to start parse data
				n := bytes.Count(buffer, []byte(I_E_NO_FLAG))
				for i := 0; i < n; i++ {

					snrinfo.WriteString("[LEVEL4_001]" + LineBreak)               // 数据还未明确
					snrinfo.WriteString("Index=" + strconv.Itoa(i+1) + LineBreak) // 数据还未明确
					snrinfo.WriteString("Value=100" + LineBreak)

					// parse info no
					i_no_s_index := bytes.Index(buffer, []byte(I_S_NO_FLAG))
					i_no_e_index := bytes.Index(buffer, []byte(I_E_NO_FLAG))
					i_no_data := buffer[i_no_s_index+len(I_S_NO_FLAG) : i_no_e_index]

					i_no_data_str := string(i_no_data) // 给后面使用

					bmpfile := filepath.Join(path, i_no_data_str+".bmp")

					// clear no data
					buffer = bytes.Replace(buffer, buffer[i_no_s_index:i_no_e_index+len(I_E_NO_FLAG)], []byte(""), 1)

					// parse info bn
					i_bn_s_index := bytes.Index(buffer, []byte(I_S_BN_FLAG))
					i_bn_e_index := bytes.Index(buffer, []byte(I_E_BN_FLAG))
					i_bn_data := buffer[i_bn_s_index+len(I_S_BN_FLAG) : i_bn_e_index]

					snrinfo.WriteString("SerialNumber=" + string(i_bn_data) + LineBreak)

					// clear bn data
					buffer = bytes.Replace(buffer, buffer[i_bn_s_index:i_bn_e_index], []byte(""), 1)

					// parse info bmp
					if i < (n - 1) { // bmpEndFlag:= "*bn{02}bn*"
						num, _ := strconv.Atoi(string(i_no_data_str))
						num = num + 1
						numstr := strconv.Itoa(num)
						if num < 10 {
							numstr = "0" + numstr
						}
						bmpEndFlag = I_S_NO_FLAG + numstr + I_E_NO_FLAG
					} else { // bmpEndFlag:= "*s[output_end]s*"
						bmpEndFlag = DATA_E_FLAG
					}
					i_bmp_s_index := bytes.Index(buffer, []byte(I_E_BN_FLAG))
					i_bmp_e_index := bytes.Index(buffer, []byte(bmpEndFlag))
					i_bmp_data := buffer[i_bmp_s_index+len(I_E_BN_FLAG) : i_bmp_e_index]

					// write bmp to file
					snrinfo.WriteString("ImageFile=" + bmpfile + LineBreak)

					file, e := os.OpenFile(bmpfile, os.O_CREATE|os.O_WRONLY, 0666)
					if e != nil {
						fmt.Println("创建bmp文件失败：" + bmpfile)
					}
					file.Write(i_bmp_data)
					if file != nil {
						file.Close()
					}

					// clear bmp data
					buffer = bytes.Replace(buffer, buffer[i_bmp_s_index:i_bmp_e_index], []byte(""), 1)

				}

				buffer = buffer[0:0]
			}

			//case <-time.After(5 * time.Second):
			//	fmt.Printf("len:%d cap:%d pointer:%p\n", len(buffer), cap(buffer), buffer)
			// 	countTimesDay += 1 // 统计交易笔数
			// 	ctd := strconv.Itoa(countTimesDay)
			// 	path := filepath.Join(config.BmpSavePath, currentDay, ctd)
			// 	os.MkdirAll(path, 0666)

			// 	for i := 1; i <= 10; i++ { // 测试写入BMP
			// 		files := filepath.Join(path, strconv.Itoa(i)+".bmp")
			// 		fmt.Println(files)
			// 		file, e := os.OpenFile(files, os.O_CREATE|os.O_WRONLY, 0666)
			// 		if e != nil {
			// 			fmt.Println(e)
			// 		}
			// 		file.Write(buffer)
			// 		file.Close()
			// 	}

		}
	}
}

func bmpClear() error {
	// 定时检查过期数据
	files, err := ioutil.ReadDir(config.BmpSavePath)
	if err != nil {
		return errors.New("未找到BMP目录：" + err.Error())
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
