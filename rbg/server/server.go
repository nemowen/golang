package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"gotest/rbg/config"
	_ "gotest/rbg/go-sql-driver/mysql"
	"gotest/rbg/logs"
	"gotest/rbg/task"
	"io/ioutil"
	//"net"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	server_preferences *config.ServerConfig // 配置文件实例
	dao                *sql.DB              // 数据库实例
	log                *logs.BeeLogger      // 日志实例
	//clientConns        map[string]*net.TCPConn // 每个客户端的连接集合
	smtp *sql.Stmt
)

// 需要传输数据的结构
type Obj struct {
	Date                string    // 日期
	Time                string    // 时间
	InTime              time.Time // 插入时间
	SerialNumber        string    // 流水号
	Type                string    // 交易类型
	CardId              string    // 预留卡号
	FaceValue           int       // 面值
	Version             int       // 版本号
	CurrencyCode        int       // 币种
	SerialNumberInTimes int       // 该钞在本笔交易内序号
	CurrencyNumber      string    // 冠字号码
	Ima                 []byte    // 冠字号图像数据
	ImaPath             string    // 冠字号保存图像路径
	ClientName          string    // 客户端设备名称
	ClientIP            string    // 客户端IP
	Remark              string    // 备注
}

// 接收数据处理方法
func (o *Obj) SendToServer(obj *Obj, replay *string) error {
	// 图像保存，当有图像地址的时候
	if "" != obj.ImaPath {
		obj.ImaPath = filepath.Join(server_preferences.BMP_SAVE_PATH, obj.ClientName, obj.Date, obj.SerialNumber+".bmp")
	REGO:
		f, err := os.Create(obj.ImaPath)
		if err != nil {
			err = os.MkdirAll(obj.ImaPath[:strings.LastIndex(obj.ImaPath, "\\")], 0666)
			if err != nil {
				log.Error("保存bmp失败：%s", obj.SerialNumber)
				*replay = config.SAVE_BMP_ERROR
				return nil
			}
			log.Info("创建目录：%s", obj.ClientName)
			goto REGO
		}
		defer f.Close()
		f.Write(obj.Ima)
	}

	// 数据存库
	//insert_sql := "INSERT INTO T_BR(SDATE,STIME,INTIME,CARDID,BILLNO,BILLBN) VALUES(?,?,?,?,?,?)"
	str_time, _ := time.Parse("2006-01-02 15:04:05", (obj.Date[0:4] + "-" + obj.Date[4:6] + "-" + obj.Date[6:8] + " " + obj.Time))
	_, err := smtp.Exec(obj.Date, obj.Time, str_time, obj.SerialNumber, obj.Type, obj.CardId, obj.FaceValue, obj.Version, obj.CurrencyCode, obj.SerialNumberInTimes, obj.CurrencyNumber, obj.ImaPath, obj.ClientName, obj.ClientIP, obj.Remark)
	if err != nil {
		log.Error("%s%s", "保存到数据库失败：", obj.CurrencyNumber)
		log.Error("%s", err)
		*replay = config.SAVE_TO_DB_ERROR
		return nil
	}
	*replay = "OK"
	return nil
}

func init() {
	loadConfig()
}

func main() {
	// 打开数据库，并连接
	openDB()

	// 以下为清理过期数据，每天22点，23点各执行一次
	dataClearTask := task.NewTask("dataClearTask", "00 59 23 * * * ", dataClear)
	task.AddTask("dataClearTask", dataClearTask)
	task.StartTask()

	// 预编译插入数据语句
	sql := "INSERT INTO T_BR(DATE,TIME,INTIME,SERIALNUMBER,TYPE,CARDID,FACEVALUE,VERSION,CURRENCYCODE,SERIALNUMBERINTIMES,BILLBN,IMAPATH,CLIENTNAME,CLIENTIP,REMARK)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	var e error
	smtp, e = dao.Prepare(sql)
	if e != nil {
		log.Error(e.Error())
	}
	defer smtp.Close()

	// 调用多CPU
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	defer dao.Close()

	//clientConns = make(map[string]*net.TCPConn, 100)

	u := new(Obj)
	rpc.Register(u)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-time.After(3 * time.Second):
	// 			for k, v := range clientConns {
	// 				log.Info("%s,%v", k, v)
	// 			}
	// 		}
	// 	}
	// }()

	log.Info("服务端已经启动! ")

	// http：方式
	exit := make(chan bool)
	rpc.HandleHTTP()
	err := http.ListenAndServe(server_preferences.SERVER_IP_PORT, nil)
	if err != nil {
		log.Error("绑定 %s 失败，请检查IP与端口是否正确! [%s]", server_preferences.SERVER_IP_PORT, err.Error())
		//exit(1)
	}

	<-exit

	// tcp 方式
	// tcpAddr, err := net.ResolveTCPAddr("tcp", server_preferences.SERVER_IP_PORT)
	// checkError(err)
	// listener, err := net.ListenTCP("tcp", tcpAddr)
	// checkError(err)

	//
	// for {
	// 	conn, err := listener.AcceptTCP()
	// 	conn.Write([]byte("ok"))
	// 	if err != nil {
	// 		log.Error("rpc.Server: accept Error:%s", err)
	// 	}
	// 	ip := strings.Split(conn.RemoteAddr().String(), ":")[0]

	// 	if v, ok := clientConns[ip]; ok {
	// 		v.Close()
	// 		continue
	// 	}
	// 	clientConns[ip] = conn
	// 	conn.SetKeepAlive(true)
	// 	conn.SetKeepAlivePeriod(120 * time.Second)
	// 	log.Info("IP [ %s ] 已经成功连接到服务器...[%d]", ip, len(clientConns))
	// 	go rpc.ServeConn(conn)
	// }

}

// 加载配置文件
func loadConfig() {
	pwd, _ := os.Getwd()
	file, e := ioutil.ReadFile(filepath.Join(pwd, "Server.Preferences.json"))
	if e != nil {
		fmt.Println("读取配置文件失败！请检查Server.Preferences.json是否在当前目录！")
		time.Sleep(30 * time.Second)
		os.Exit(1)
	}

	// 日志文件
	logfile := filepath.Join(pwd, "logs", "server.log")
	os.MkdirAll(logfile[0:len(logfile)-10], 0666)
	log = logs.NewLogger(100000)
	_, e = os.Stat(logfile)
	if nil != e {
		os.Create(logfile)
	}
	log.SetLogger("file", `{"filename":"`+strings.Replace(logfile, "\\", "/", -1)+`"}`)
	log.SetLogger("console", "")
	server_preferences = new(config.ServerConfig)
	e = json.Unmarshal(file, server_preferences)
	if e != nil {
		fmt.Println("解析配置文件失败!" + e.Error())
		//os.Exit(2)
	}
}

// 获取数据库
func openDB() {

	// 检查数据库连接是否有误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			time.Sleep(30 * time.Second)
			os.Exit(1)
		}
	}()

	// 打开数据库
	db, err := sql.Open("mysql", server_preferences.DATABASE_USER_NAME+":"+server_preferences.DATABASE_PASSWORD+"@tcp(127.0.0.1:3306)/"+server_preferences.DATABASE_NAME+"?charset=utf8") // &timeout=60s
	if err != nil {
		panic("错误：打开数据库失败: " + err.Error())
	}
	// 测试连接数据库
	err = db.Ping()
	if err != nil {
		panic("错误：连接数据库失败:" + err.Error())
	}
	// 设置数据库最大空间连接数
	db.SetMaxIdleConns(server_preferences.DB_MAX_IDLE_CONNS)
	// 设置数据库最大连接数
	db.SetMaxOpenConns(server_preferences.DB_MAX_OPEN_CONNS)
	dao = db

}

// 关闭客户端连接
func CloseConn() {
	if nil != dao {
		dao.Close()
		dao = nil
	}
}

// 清理过期数据
func dataClear() error {
	clearDataDay := time.Now().Add(time.Hour * -(server_preferences.DATA_KEEPING_DAYS * 24))
	_, err := dao.Exec("DELETE FROM T_BR WHERE INTIME <= ? ", clearDataDay)
	if err != nil {
		log.Warn("删除数据库过期数据失败: %s", err.Error())
		return errors.New("删除数据库过期数据失败：" + err.Error())
	} else {
		log.Info("删除数据库过期数据成功!")
	}
	files, err := ioutil.ReadDir(server_preferences.BMP_SAVE_PATH)
	if err != nil {
		log.Warn("未找到BMP目录: %s", err.Error())
		return errors.New("未找到BMP目录：" + err.Error())
	}
	clearDataDayInt, _ := strconv.Atoi(clearDataDay.Format("20060102"))
	for _, file := range files {
		// 读取每个客户端BMP目录
		dayBmpfiles, _ := ioutil.ReadDir(filepath.Join(server_preferences.BMP_SAVE_PATH, file.Name()))
		for _, bmpDir := range dayBmpfiles {
			bmpDirName, _ := strconv.Atoi(bmpDir.Name())
			if bmpDirName <= clearDataDayInt {
				err = os.RemoveAll(filepath.Join(server_preferences.BMP_SAVE_PATH, file.Name(), bmpDir.Name()))
				if err != nil {
					log.Warn("删除过期BMP失败: %s", err.Error())
					return errors.New("删除过期BMP失败:" + err.Error())
				} else {
					log.Info("客户端:[%s] %s bmp数据已经清理", file.Name(), bmpDir.Name())
				}
			}
		}
	}
	log.Info("成功清理%d天以前的数据！", server_preferences.DATA_KEEPING_DAYS)
	return nil
}
