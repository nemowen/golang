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
}

var server_preferences config.ServerConfig
var db *sql.DB

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

	//打开数据库
	db = OpenDatabase()
}

func OpenDatabase() *sql.DB {
	db, err := sql.Open("mysql", server_preferences.DATABASE_NAME+":"+server_preferences.DATABASE_PASSWORD+"@tcp("+server_preferences.DATABASE_IP+":"+server_preferences.DATABASE_PORT+")/"+server_preferences.DATABASE_NAME+"?charset=utf8")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	return db
}

var w sync.WaitGroup

func (o *Obj) SendToServer(obj *Obj, replay *string) error {
	saveObj(obj)
	*replay = "ok"
	return nil
}

func saveObj(obj *Obj) {
	w.Add(2)
	go func() {
		s := obj.CurrencyNumber + "|" + obj.Date + "|" + string(obj.ID) + "|" + obj.ImaPath + "|" + obj.Type + "|" + string(obj.SerialNumberInTimes)
		saveTo, e := os.Create(server_preferences.BMP_SAVE_PATH + obj.ID + ".txt")
		if e != nil {
			log.Println(e)
		}
		defer saveTo.Close()
		_, err := saveTo.WriteString(s)
		if err != nil {
			log.Println("保存信息失败", obj.ID, err)
		}
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
