package rpcobj

import (
	"gotest/rbg/config"
	"gotest/rbg/server/db"
	"gotest/rbg/server/utils"
	"log"
	"os"
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

func (o *Obj) SendToServer(obj *Obj, replay *string) error {
	insert_sql := "INSERT INTO T_BR(SDATE,STIME,INTIME,CARDID,BILLNO,BILLBN) VALUES(?,?,?,?,?,?)"
	str_time, _ := time.Parse("2006-01-02 15:04:05", (obj.Date[0:4] + "-" + obj.Date[4:6] + "-" + obj.Date[6:8] + " " + obj.Time))

	_, err := db.Dao.Exec(insert_sql, obj.Date, obj.Time, str_time, obj.CardId, obj.SerialNumberInTimes, obj.CurrencyNumber)
	if err != nil {
		log.Println("保存到数据库失败：", obj.CurrencyNumber)
		*replay = config.SAVE_TO_DB_ERROR
		return nil
	}

	f, err := os.Create(utils.Server_preferences.BMP_SAVE_PATH + obj.SerialNumber + ".bmp")
	if err != nil {
		log.Println("保存bmp失败：", obj.SerialNumber)
		*replay = config.SAVE_BMP_ERROR
		return nil
	}
	defer f.Close()
	f.Write(obj.Ima)

	*replay = "OK"
	return nil
}
