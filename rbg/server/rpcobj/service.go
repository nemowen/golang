package rpcobj

import (
	"database/sql"
	"gotest/rbg/config"
	"gotest/rbg/server/db"
	"gotest/rbg/server/utils"
	"log"
	"os"
	"sync"
	"time"
)

var dao *sql.DB
var server_preferences *config.ServerConfig

//同步工作组机制
var w sync.WaitGroup

func init() {
	dao = db.OpenDB()
	server_preferences = utils.LoadConfig()
}

func saveObj(obj *Obj) {
	w.Add(2)
	go func() {
		insert_sql := "INSERT INTO T_BR(SDATE,STIME,INTIME,CARDID,BILLNO,BILLBN) VALUES(?,?,?,?,?,?)"
		str_time, _ := time.Parse("2006-01-02 15:04:05", (obj.Date + " " + obj.Time))
		dao.Exec(insert_sql, obj.Date, obj.Time, str_time, obj.CardId, obj.SerialNumberInTimes, obj.CurrencyNumber)
		w.Done()
	}()

	go func() {
		f, err := os.Create(server_preferences.BMP_SAVE_PATH + obj.SerialNumber + ".bmp")
		if err != nil {
			log.Println("保存bmp失败：", obj.SerialNumber)
		}
		f.Write(obj.Ima)
		f.Close()
		w.Done()
	}()
	w.Wait()
}
