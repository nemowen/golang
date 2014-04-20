package main

import (
	"bufio"
	"gotest/rbg/server/rpcobj"

	"log"
	"net/rpc"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	ADDR                  string        = "192.168.0.109:1314" //IP地址与端口
	WATCHE_FILE           string        = "D:/EN/flag.ini"     //监控文件
	NOTE_FILE             string        = "D:/EN/note.ini"     //note文件
	START_WORK_FLAG       string        = "OK"                 //开始工作状态
	FLAG_STATE_SPEED      time.Duration = 5                    //读取flag.ini文件速度， 称为单位
	RECONNECT_SERVER_TIME time.Duration = 30                   //当连接失败时，多少秒后重新连接服务器
)

var (
	client    *rpc.Client //连接服务器client实例
	t         time.Time
	reply     chan string
	read      chan bool
	noteBufer *bufio.Reader //需要读取的文件
)

func init() {
	client = connect()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 用户接收监控文件的状态
	reply = make(chan string, 1)
	read = make(chan bool)

	// 启动监控文件
	go NewWatcher(WATCHE_FILE, reply, read)
	read <- true

	// 读取传回的监控内容
	for d := range reply {
		if strings.Contains(strings.ToUpper(d), START_WORK_FLAG) {
			read <- false
			// 开始采集并发送数据
			sendDataToServer()
			read <- true
		} else {
			read <- true
			continue
		}
	}

	defer closeConn(client)
}

func done() {
	f, e := os.Create(WATCHE_FILE)
	if e != nil {
		log.Println("打开文件失败：", WATCHE_FILE)
		return
	}
	defer f.Close()
	// 将Flag标识位置为END
	f.WriteString("END")
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
			defer f.Close()
			buf := make([]byte, 8)
			f.Read(buf)
			reply <- string(buf)
		}
		time.Sleep(FLAG_STATE_SPEED * time.Second)
	}
}

func connect() (client *rpc.Client) {
	for client == nil {
		var e error
		client, e = rpc.DialHTTP("tcp", ADDR)
		if e != nil {
			log.Println("连接服务器失败,请检查网络或服务器是否已经启动...")
			log.Println("30秒后自动重新连接...")
			time.Sleep(RECONNECT_SERVER_TIME * time.Second)
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

func sendDataToServer() {
	t = time.Now()
	//获取note文件
	f, e := os.Open(NOTE_FILE)
	if e != nil {
		log.Println("打开文件失败：", NOTE_FILE)
	}
	defer f.Close()
	noteBufer = bufio.NewReader(f)
	var line string
	var err error
	for err == nil {
		line, err = noteBufer.ReadString('\n')
		items := strings.Split(line, "|")
		obj := new(rpcobj.Obj)
		obj.Date = items[0]
		obj.Time = items[1]
		obj.ID = items[2]
		obj.Type = items[3]
		obj.FaceValue, _ = strconv.Atoi(items[5])
		obj.Version, _ = strconv.Atoi(items[6])
		obj.SerialNumberInTimes, _ = strconv.Atoi(items[7])
		obj.Number = items[8]
		f, _ := os.Open(items[9])

		defer f.Close()
		b := make([]byte, 5<<10)
		f.Read(b)

		obj.Ima = b

		replay := new(string)
		// call method: 同步方式，似乎效率比go 异步方法高
		err := client.Call("Obj.SendToServer", obj, replay)
		if err != nil {
			log.Println("失去连接...")
			closeConn(client)
			log.Println("正重新连接中...")
			client = connect()

		}

		log.Println(">>>>>>", obj.ID, *replay)
	}
	log.Println("已经完成本笔任务，用时：", time.Now().Sub(t))

	done()
}
