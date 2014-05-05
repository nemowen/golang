package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gotest/rbg/server/utils"
	"log"
)

const (
	_DB_MAX_IDLE_CONNS int = 100
	_DB_MAX_OPEN_CONNS int = 1000
)

var dao *sql.DB

// 获取数据库
func OpenDB() *sql.DB {
	if dao != nil {
		return dao
	}
	config := utils.LoadConfig()
	dao, err := sql.Open("mysql", config.DATABASE_USER_NAME+":"+config.DATABASE_PASSWORD+"@/"+config.DATABASE_NAME+"?charset=utf8")
	if err != nil {
		errmsg := "错误：连接数据库连接失败！"
		log.Println(errmsg)
		log.Println(err)
	}
	dao.SetMaxIdleConns(_DB_MAX_IDLE_CONNS)
	dao.SetMaxOpenConns(_DB_MAX_OPEN_CONNS)
	log.Println("连接数据库成功!！")
	return dao
}

// 关闭客户端连接
func CloseConn() {
	if nil != dao {
		dao.Close()
		dao = nil
	}
}
