package rpcobj

import (
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
	SerialNumberInTimes int       //该钞在本笔交易内序号
	CurrencyNumber      string    //冠字号码
	Ima                 []byte    //冠字号图像数据
	ImaPath             string    //冠字号保存图像路径

	ClientName string //客户端设备名称
	ClientIP   string //客户端IP
}

func (o *Obj) SendToServer(obj *Obj, replay *string) error {
	saveObj(obj)
	*replay = "ok"
	return nil
}
