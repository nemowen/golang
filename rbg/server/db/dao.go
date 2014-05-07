package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gotest/rbg/server/utils"
	"log"
	"os"
	"sync"
)

var Dao *sql.DB
var once sync.Once

func init() {
	once.Do(openDB)
}

// 获取数据库
func openDB() {
	config := utils.Server_preferences
	db, err := sql.Open("mysql", config.DATABASE_USER_NAME+":"+config.DATABASE_PASSWORD+"@/"+config.DATABASE_NAME+"?charset=utf8")
	if err != nil {
		errmsg := "错误：连接数据库连接失败！"
		log.Println(errmsg)
		log.Println(err)
		os.Exit(1)
	}
	db.SetMaxIdleConns(config.DB_MAX_IDLE_CONNS)
	db.SetMaxOpenConns(config.DB_MAX_OPEN_CONNS)
	log.Println("连接数据库成功!！")
	Dao = db
}

// 关闭客户端连接
func CloseConn() {
	if nil != Dao {
		Dao.Close()
		Dao = nil
	}
}
