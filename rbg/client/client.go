package main

import (
	"bufio"
	//"fmt"
	"gotest/rbg/server/rpcobj"
	"log"
	//"net"
	"net/rpc"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	ADDR                  string        = "127.0.0.1:1314" //IP地址与端口
	WATCHE_FILE           string        = "D:/EN/flag.ini" //监控文件
	NOTE_FILE             string        = "D:/EN/note.ini" //note文件
	START_WORK_FLAG       string        = "OK"             //开始工作状态
	FLAG_STATE_SPEED      time.Duration = 5                //读取flag.ini文件速度， 称为单位
	RECONNECT_SERVER_TIME time.Duration = 30               //当连接失败时，多少秒后重新连接服务器
)

var (
	client    *rpc.Client //连接服务器client实例
	t         time.Time
	reply     chan string
	read      chan bool
	noteBufer *bufio.Reader //需要读取的文件
	rebackObj *rpcobj.Obj   //当网络在传输过程中失败时，回滚的对象
	obj       *rpcobj.Obj   //需要传输的对象
	//handwareAddrs map[string]string
)

func init() {
	// 以下读取网卡信息
	// Interface, err := net.Interfaces()
	// if err != nil {
	// 	panic("未发现网卡地址")
	// 	os.Exit(1)
	// }
	// handwareAddrs = make(map[string]string, len(Interface))
	// for _, inter := range Interface {
	// 	inMAC := strings.ToUpper(inter.HardwareAddr.String())
	// 	handwareAddrs[inMAC] = inMAC
	// }

	// if len(os.Args) != 2 {
	// 	fmt.Println("为保障安全:请先绑定本机上的网卡地址")
	// 	os.Exit(0)
	// }

	// addr := os.Args[1]
	// h, e := net.ParseMAC(addr)
	// if e != nil {
	// 	fmt.Println("为保障安全:请先绑定本机上的网卡地址")
	// 	fmt.Println("方法：client.exe 90-4C-E5-58-7E-FE")
	// 	os.Exit(2)
	// }
	// inputMAC := strings.ToUpper(h.String())
	// if inputMAC != handwareAddrs[inputMAC] {
	// 	fmt.Println("网卡地址不匹配")
	// 	os.Exit(0)
	// }

	//开始连接服务器
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

	//清空数据
	note, e := os.Create(NOTE_FILE)
	if e != nil {
		log.Println("打开文件失败：", NOTE_FILE)
		return
	}
	defer note.Close()
	//note.WriteString("")
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
			log.Println("连接服务器失败,请检查网络或服务器是否启动...")
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
		//如果回滚对象为空，正常运行，否则先处理上次失败的对象
		if rebackObj == nil {
			line, err = noteBufer.ReadString('\n')
			if 10 > len(line) {
				log.Println("数据有误:", line)
				break
			}
			items := strings.Split(line, "|")
			obj = new(rpcobj.Obj)
			obj.Date = items[0]
			obj.Time = items[1]
			obj.ID = items[2]
			obj.Type = items[3]
			obj.FaceValue, _ = strconv.Atoi(items[5])
			obj.Version, _ = strconv.Atoi(items[6])
			obj.SerialNumberInTimes, _ = strconv.Atoi(items[7])
			obj.Number = items[8]

			//读取图像数据
			f, _ := os.Open(items[9])
			defer f.Close()
			b := make([]byte, 5<<10)
			f.Read(b)
			obj.Ima = b
		} else {
			obj = rebackObj
		}

		replay := new(string)

		//log.Println(">>>>>> 正在传输", obj.Number)

		//为制造网络断开，加时，模拟网络异常
		//time.Sleep(3 * time.Second)

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

		log.Println(">>>>>> 上传成功", obj.Number, *replay)
		// 清除回滚对象
		rebackObj = nil

	}
	log.Println("已经完成本笔任务，用时：", time.Now().Sub(t))

	done()
}
