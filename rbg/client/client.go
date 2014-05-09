package main

import (
	"bufio"
	"encoding/json"
	"gotest/rbg/config"
	"gotest/rbg/logs"
	"io/ioutil"
	"net"
	"net/rpc"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Obj struct {
	Date                string    //日期
	Time                string    //时间
	InTime              time.Time //插入时间
	SerialNumber        string    //流水号
	Type                string    //交易类型
	CardId              string    //预留卡号
	FaceValue           int       //面值
	Version             int       //版本号
	CurrencyCode        int       //币种
	SerialNumberInTimes int       //该钞在本笔交易内序号
	CurrencyNumber      string    //冠字号码
	Ima                 []byte    //冠字号图像数据
	ImaPath             string    //冠字号保存图像路径

	ClientName string //客户端设备名称
	ClientIP   string //客户端IP
}

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
	rebackObj *Obj
	// 需要传输的对象
	obj *Obj
	// 配置文件
	client_preferences config.ClientConfig
	//日志记录
	log *logs.BeeLogger
)

func init() {
	//加载配置文件
	file, e := ioutil.ReadFile(CLIENT_PREFERENCES)
	if e != nil {
		panic("读取配置文件失败！请与管理员联系！")
		os.Exit(1)
	}
	json.Unmarshal(file, &client_preferences)

	//日志初始化
	log = logs.NewLogger(10000)
	//日志文件记录
	log.SetLogger("file", `{"filename":"`+client_preferences.LOG_SAVE_PATH+`"}`)
	//日志终端记录
	log.SetLogger("console", "")
	//log.SetLogger("smtp", `{"username":"nemo.emails@gmail.com","password":"","host":"smtp.gmail.com:587","sendTos":["wenbin171@163.com"],"level":4}`)

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

//连接服务器
func connect() (client *rpc.Client) {
	for client == nil {
		var e error
		client, e = rpc.Dial("tcp", client_preferences.SERVER_IP_PORT)
		if e != nil {
			log.Error("连接服务器失败,请检查网络或服务器是否启动...")
			log.Info("%d%s", client_preferences.RECONNECT_SERVER_TIME.Nanoseconds(), "秒后自动重新连接...")
			time.Sleep(client_preferences.RECONNECT_SERVER_TIME * time.Second)
			continue
		} else {
			log.Info("连接服务器成功...")
			break
		}
	}
	return client
}

// 监控flag.ini文件状态
func NewWatcher(filepath string, reply chan string, read chan bool) {
	for {
		if ok := <-read; ok {
			f, e := os.Open(filepath)
			if e != nil {
				log.Warn("打开文件失败：%s", filepath)
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

//收集数据到服务器
func sendDataToServer() {
	log.Info("正在上传数据...")
	t = time.Now()
	//获取note文件
	f, e := os.Open(client_preferences.NOTE_FILE_PATH)
	if e != nil {
		log.Warn("打开文件失败：%s", client_preferences.NOTE_FILE_PATH)
	}
	noteBufer = bufio.NewReader(f)
	defer closeFile(f)

	var line string
	var err error
	replay := new(string)
	for err == nil {
		//如果回滚对象为空，正常运行，否则先处理上次失败的对象
		if rebackObj == nil {
			line, err = noteBufer.ReadString('\n')
			if 10 > len(line) {
				log.Warn("数据有误:%s", line)
				continue
			}
			line = strings.TrimRight(line, "\r\n")
			items := strings.Split(line, "|")
			obj = new(Obj)
			obj.Date = items[0]
			obj.Time = items[1]
			obj.SerialNumber = items[2]
			obj.Type = items[3]
			obj.CardId = items[4]
			obj.FaceValue, _ = strconv.Atoi(items[5])
			obj.Version, _ = strconv.Atoi(items[6])
			obj.CurrencyCode, _ = strconv.Atoi(items[7])
			obj.SerialNumberInTimes, _ = strconv.Atoi(items[8])
			obj.CurrencyNumber = items[9]
			obj.ImaPath = items[10]

			//读取图像数据
			f, e := os.Open(obj.ImaPath)
			if e != nil {
				log.Error("读取图像数据失败：[%s] (%s)", obj.SerialNumber, e)
			}
			b, _ := ioutil.ReadAll(f)
			f.Read(b)
			closeFile(f)

			obj.Ima = b
		} else {
			obj = rebackObj
		}

		// call method: 同步方式，似乎效率比go 异步方法高
		err := client.Call("Obj.SendToServer", obj, replay)
		if err != nil {
			//出现网络中断时，回滚并保存当前对象
			rebackObj = obj
			log.Error("与服务器失去连接...")
			closeConn(client)
			log.Info("正重新连接中...")
			client = connect()
			continue
		}

		log.Trace("[%s] STATE:%s", obj.CurrencyNumber, *replay)

		if strings.Contains(config.SAVE_TO_DB_ERROR, *replay) {
			log.Warn(">>>>>> 服务器保存数据失败，10秒后重新上传:%s", obj.CurrencyNumber)
			rebackObj = obj
			time.Sleep(10 * time.Second)
			continue
		}
		// 清除回滚对象
		rebackObj = nil

	}
	log.Info("已经完成本笔任务，用时：%s", time.Now().Sub(t).String())

	// 完成本次任务后续操作
	done()
}

//完成一笔数据后续操作
func done() {
	f, e := os.Create(client_preferences.FLAG_FILE_PATH)
	if e != nil {
		log.Info("打开文件失败：", client_preferences.FLAG_FILE_PATH)
		return
	}
	// 将Flag标识位置为END
	f.WriteString("END")
	defer f.Close()

	//清空数据
	note, e := os.Create(client_preferences.NOTE_FILE_PATH)
	if e != nil {
		log.Info("打开文件失败：", client_preferences.NOTE_FILE_PATH)
		return
	}
	defer note.Close()
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

func getLocalIPAddr() string {
	var ip string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Warn(err.Error())
	}
	for _, addr := range addrs {
		ips := addr.String()
		if "0.0.0.0" != ips {
			ip = ips
			break
		}
	}
	return ip
}

func getPublicIPAddr() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		log.Warn(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
