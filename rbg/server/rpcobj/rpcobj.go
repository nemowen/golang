package rpcobj

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"gotest/rbg/config"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

const (
	OK    int = 200
	ERROR int = 100
)

// 需要传输数据的结构
type Obj struct {
	Date                string    //日期
	Time                string    //时间
	InTime              time.Time //插入时间
	ID                  string    //流水号
	Type                string    //交易类型
	CardId              string    //预留卡号
	FaceValue           int       //面值
	Version             int       //版本号
	SerialNumberInTimes int       //该钞在本笔交易内序号
	CurrencyNumber      string    //冠字号码
	Ima                 []byte    //冠字号图像数据
	ImaPath             string    //冠字号保存图像路径

	ClientName string //客户端设备名称
	ClientIP   string //客户端IP
}

var server_preferences config.ServerConfig

//服务器数据库连接池，一个客户端分配一个连接
//var conns map[string]*sql.DB
var db *sql.DB

//同步工作组机制
var w sync.WaitGroup

const (
	//客户端配置文件路径
	SERVER_PREFERENCES string = "C:/Windows/Server.Preferences.json"
)

func init() {
	//加载配置文件
	file, e := ioutil.ReadFile(SERVER_PREFERENCES)
	if e != nil {
		panic("读取配置文件失败！请与管理员联系！")
		os.Exit(1)
	}

	json.Unmarshal(file, &server_preferences)

	//分配连接池大小
	//conns = make(map[string]*sql.DB)

	db = GetConn()
}

// 获取连接
func GetConn() *sql.DB {

	db, err := sql.Open("mysql", server_preferences.DATABASE_USER_NAME+":"+server_preferences.DATABASE_PASSWORD+"@/"+server_preferences.DATABASE_NAME+"?charset=utf8")
	if err != nil {
		errmsg := "错误：创建数据库连接失败!"
		log.Println(errmsg)
		log.Println(err)
		os.Exit(1)
	}
	return db
}

// 关闭客户端连接
// func (c *Obj) CloseConn(clientName string, replay *int) error {
// 	dbo := conns[clientName]
// 	dbo.Close()
// 	delete(conns, clientName)

// 	*replay = OK
// 	return nil
// }

func (o *Obj) SendToServer(obj *Obj, replay *string) error {
	saveObj(obj)
	*replay = "ok"
	return nil
}

func saveObj(obj *Obj) {
	w.Add(2)
	go func() {

		insert_sql := "INSERT INTO T_BR(SDATE,STIME,INTIME,CARDID,BILLNO,BILLBN) VALUES(?,?,?,?,?,?)"
		str_time, _ := time.Parse("2006-01-02 15:04:05", obj.Date+" "+obj.Time)
		_, e4 := db.Exec(insert_sql, obj.Date, obj.Time, str_time, obj.CardId, obj.ID, obj.CurrencyNumber)
		log.Println(e4)

		//s := obj.CurrencyNumber + "|" + obj.Date + "|" + string(obj.ID) + "|" + obj.ImaPath + "|" + obj.Type + "|" + string(obj.SerialNumberInTimes)
		//saveTo, e := os.Create(server_preferences.BMP_SAVE_PATH + obj.ID + ".txt")
		//if e != nil {
		//	log.Println(e)
		//}
		//defer saveTo.Close()
		//_, err := saveTo.WriteString(s)
		//if err != nil {
		//	log.Println("保存信息失败", obj.ID, err)
		//}
		w.Done()
	}()

	go func() {
		f, err := os.Create(server_preferences.BMP_SAVE_PATH + obj.ID + ".bmp")
		if err != nil {
			log.Println("保存bmp失败：", obj.ID)
		}
		f.Write(obj.Ima)
		f.Close()
		w.Done()
	}()
	w.Wait()
}
