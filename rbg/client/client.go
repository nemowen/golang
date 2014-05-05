package main

import (
	"bufio"
	"encoding/json"
	"gotest/rbg/config"
	"gotest/rbg/server/rpcobj"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	//客户端配置文件路径
	CLIENT_PREFERENCES string = "C:/Windows/Client.Preferences.json"
)

var (
	// 连接服务器client实例
	client *rpc.Client
	// 采集用时统计
	t time.Time
	// 用户接收监控文件的状态
	reply chan string
	// 是否可以开始采集数据
	read chan bool
	// 需要读取的文件
	noteBufer *bufio.Reader
	// 当网络在传输过程中失败时，回滚的对象
	rebackObj *rpcobj.Obj
	// 需要传输的对象
	obj *rpcobj.Obj
	// 配置文件
	client_preferences config.ClientConfig
)

func init() {
	//加载配置文件
	file, e := ioutil.ReadFile(CLIENT_PREFERENCES)
	if e != nil {
		panic("读取配置文件失败！请与管理员联系！")
		os.Exit(1)
	}
	json.Unmarshal(file, &client_preferences)

	//开始连接服务器
	client = connect()
}

func main() {
	// 启动CPU运行数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	reply = make(chan string, 1)
	read = make(chan bool)
	// 启动监控文件
	go NewWatcher(client_preferences.FLAG_FILE_PATH, reply, read)
	read <- true

	// 读取传回的监控内容
	for d := range reply {
		if strings.Contains(strings.ToUpper(d), client_preferences.START_WORK_FLAG) {
			read <- false
			// 开始采集并发送数据
			sendDataToServer()
			read <- true
		} else {
			read <- true
			continue
		}
	}

	// 关闭连接
	defer closeConn(client)
}

func done() {
	f, e := os.Create(client_preferences.FLAG_FILE_PATH)
	if e != nil {
		log.Println("打开文件失败：", client_preferences.FLAG_FILE_PATH)
		return
	}
	// 将Flag标识位置为END
	f.WriteString("END")
	defer f.Close()

	//清空数据
	note, e := os.Create(client_preferences.NOTE_FILE_PATH)
	if e != nil {
		log.Println("打开文件失败：", client_preferences.NOTE_FILE_PATH)
		return
	}
	defer note.Close()
}

// 监控flag.ini文件状态
func NewWatcher(filepath string, reply chan string, read chan bool) {
	for {
		if ok := <-read; ok {
			f, e := os.Open(filepath)
			if e != nil {
				log.Println("打开文件失败：", filepath)
				continue
			}
			buf, _ := ioutil.ReadAll(f)
			closeFile(f)
			reply <- string(buf)
		}
		// 检查频率不能小于2秒
		if 2 > client_preferences.FLAG_STATE_SPEED {
			client_preferences.FLAG_STATE_SPEED = 2
		}
		time.Sleep(client_preferences.FLAG_STATE_SPEED * time.Second)

	}
}

func connect() (client *rpc.Client) {
	for client == nil {
		var e error
		client, e = rpc.DialHTTP("tcp", client_preferences.SERVER_IP_PORT)
		if e != nil {
			log.Println("连接服务器失败,请检查网络或服务器是否启动...")
			log.Println(client_preferences.RECONNECT_SERVER_TIME.Nanoseconds(), "秒后自动重新连接...")
			time.Sleep(client_preferences.RECONNECT_SERVER_TIME * time.Second)
			continue
		} else {
			log.Println("连接服务器成功...")
			break
		}
	}
	return client
}

func closeConn(client *rpc.Client) {
	if client != nil {
		client.Close()
	}
}

func closeFile(f *os.File) {
	if f != nil {
		f.Close()
		f = nil
	}
}

func sendDataToServer() {
	// var states int
	// errs := client.Call("Obj.GetConn", obj.ClientName, &states)
	// if errs != nil {
	// 	log.Println(errs)
	// }

	t = time.Now()
	//获取note文件
	f, e := os.Open(client_preferences.NOTE_FILE_PATH)
	if e != nil {
		log.Println("打开文件失败：", client_preferences.NOTE_FILE_PATH)
	}
	noteBufer = bufio.NewReader(f)
	defer closeFile(f)

	var line string
	var err error
	for err == nil {
		//如果回滚对象为空，正常运行，否则先处理上次失败的对象
		if rebackObj == nil {
			line, err = noteBufer.ReadString('\n')
			if 10 > len(line) {
				log.Println("数据有误:", line)
				break
			}
			line = strings.TrimRight(line, "\r\n")
			items := strings.Split(line, "|")
			obj = new(rpcobj.Obj)
			obj.Date = items[0]
			obj.Time = items[1]
			obj.SerialNumber = items[2]
			obj.Type = items[3]
			obj.CardId = items[4]
			obj.FaceValue, _ = strconv.Atoi(items[5])
			obj.Version, _ = strconv.Atoi(items[6])
			obj.SerialNumberInTimes, _ = strconv.Atoi(items[7])
			obj.CurrencyNumber = items[9]
			obj.ImaPath = items[10]

			//读取图像数据
			f, e := os.Open(obj.ImaPath)
			if e != nil {
				log.Println("读取图像数据失败：["+obj.CurrencyNumber+"]"+"["+obj.SerialNumber+"]", e)
			}
			b, _ := ioutil.ReadAll(f)
			f.Read(b)
			closeFile(f)

			obj.Ima = b
		} else {
			obj = rebackObj
		}

		replay := new(string)

		// call method: 同步方式，似乎效率比go 异步方法高
		err := client.Call("Obj.SendToServer", obj, replay)
		if err != nil {
			//出现网络中断时，回滚并保存当前对象
			rebackObj = obj
			log.Println("与服务器失去连接...")
			closeConn(client)
			log.Println("正重新连接中...")
			client = connect()
			continue
		}
		log.Println(">>>>>> 上传成功", obj.CurrencyNumber, *replay)
		// 清除回滚对象
		rebackObj = nil

	}
	log.Println("已经完成本笔任务，用时：", time.Now().Sub(t))

	// 完成本次任务后续操作
	done()
}
