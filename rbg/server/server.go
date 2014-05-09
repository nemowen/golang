package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gotest/rbg/config"
	"gotest/rbg/logs"
	"io/ioutil"
	"net"
	"net/rpc"
	"os"
	"runtime"
	"time"
)

// 需要传输数据的结构
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

// 接收数据处理方法
func (o *Obj) SendToServer(obj *Obj, replay *string) error {
	insert_sql := "INSERT INTO T_BR(SDATE,STIME,INTIME,CARDID,BILLNO,BILLBN) VALUES(?,?,?,?,?,?)"
	str_time, _ := time.Parse("2006-01-02 15:04:05", (obj.Date[0:4] + "-" + obj.Date[4:6] + "-" + obj.Date[6:8] + " " + obj.Time))
	_, err := dao.Exec(insert_sql, obj.Date, obj.Time, str_time, obj.CardId, obj.SerialNumberInTimes, obj.CurrencyNumber)
	if err != nil {
		log.Error("%s%s", "保存到数据库失败：", obj.CurrencyNumber)
		log.Error("%s", err)
		*replay = config.SAVE_TO_DB_ERROR
		return nil
	}

	f, err := os.Create(server_preferences.BMP_SAVE_PATH + obj.SerialNumber + ".bmp")
	if err != nil {
		log.Error("保存bmp失败：", obj.SerialNumber)
		*replay = config.SAVE_BMP_ERROR
		return nil
	}
	defer f.Close()
	f.Write(obj.Ima)

	*replay = "OK"
	return nil
}

const (
	//客户端配置文件路径
	_server_preferences string = "C:/Windows/Server.Preferences.json"
)

var server_preferences *config.ServerConfig
var dao *sql.DB
var log *logs.BeeLogger

func init() {
	loadConfig()

	log = logs.NewLogger(1000)

	log.SetLogger("file", `{"filename":"`+server_preferences.LOG_SAVE_PATH+`"}`)
	log.SetLogger("console", "")
	//log.SetLogger("smtp", `{"username":"nemo.emails@gmail.com","password":"'sytwgmail%100s.","host":"smtp.gmail.com:587","sendTos":["wenbin171@163.com"],"level":4}`)

	openDB()
}

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	u := new(Obj)
	rpc.Register(u)

	tcpAddr, err := net.ResolveTCPAddr("tcp", server_preferences.SERVER_IP_PORT)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	log.Info("服务已经启动!")

	//rpc.Accept(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("rpc.Server: accept Error:%s", err)
		}
		log.Info(" IP [ %s ] 已经连接到服务器...", conn.RemoteAddr().String())
		go rpc.ServeConn(conn)
	}

}

//加载配置文件
func loadConfig() {
	server_preferences = new(config.ServerConfig)
	file, e := ioutil.ReadFile(_server_preferences)
	if e != nil {
		fmt.Println("读取配置文件失败!请与管理员联系!")
		os.Exit(1)
	}
	json.Unmarshal(file, server_preferences)
}

// 获取数据库
func openDB() {
	config := server_preferences
	db, err := sql.Open("mysql", config.DATABASE_USER_NAME+":"+config.DATABASE_PASSWORD+"@/"+config.DATABASE_NAME+"?charset=utf8")
	if err != nil {
		errmsg := "错误：连接数据库连接失败!"
		fmt.Println(errmsg)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	db.SetMaxIdleConns(config.DB_MAX_IDLE_CONNS)
	db.SetMaxOpenConns(config.DB_MAX_OPEN_CONNS)
	dao = db
}

// 关闭客户端连接
func CloseConn() {
	if nil != dao {
		dao.Close()
		dao = nil
	}
}

func checkError(err error) {
	if err != nil {
		log.Warn("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
