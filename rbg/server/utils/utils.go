package utils

import (
	"encoding/json"
	"gotest/rbg/config"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const (
	//客户端配置文件路径
	_Server_preferences string = "C:/Windows/Server.Preferences.json"
)

var Server_preferences *config.ServerConfig
var once sync.Once

func init() {
	once.Do(loadConfig)
}

//加载配置文件
func loadConfig() {
	Server_preferences = new(config.ServerConfig)
	file, e := ioutil.ReadFile(_Server_preferences)
	if e != nil {
		panic("读取配置文件失败！请与管理员联系！")
		os.Exit(1)
	}
	json.Unmarshal(file, Server_preferences)
	log.Println("加载配置文件成功！")
}
