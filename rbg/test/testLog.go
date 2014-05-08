package main

import (
	"gotest/rbg/logs"
	//"runtime"
	"time"
)

func main() {

	log := logs.NewLogger(1000)
	path := "D:/test.log"
	log.SetLogger("file", `{"filename":"`+path+`"}`)
	log.SetLogger("console", "")
	log.SetLogger("smtp", `{"username":"nemo.emails@gmail.com","password":"","host":"smtp.gmail.com:587","sendTos":["wenbin171@163.com"],"level":4}`)

	log.Info("%s", "test........")
	log.Error("%s", "test........")

	log.Trace("trace %s %s", "param1", "param2")
	log.Debug("debug")
	log.Info("info")
	log.Warn("warning")
	log.Error("error")
	log.Critical("critical")

	//log.Flush()
	time.Sleep(10 * time.Second)

	//log.Close()

}
