package utils

import (
	"encoding/json"
	"gotest/rbg/config"
	"io/ioutil"
	"log"
	"os"
)

const (
	//客户端配置文件路径
	_SERVER_PREFERENCES string = "C:/Windows/Server.Preferences.json"
)

var server_preferences *config.ServerConfig

//加载配置文件
func LoadConfig() *config.ServerConfig {
	if server_preferences != nil {
		return server_preferences
	}
	server_preferences = new(config.ServerConfig)
	file, e := ioutil.ReadFile(_SERVER_PREFERENCES)
	if e != nil {
		panic("读取配置文件失败！请与管理员联系！")
		os.Exit(1)
	}
	json.Unmarshal(file, server_preferences)
	log.Println("加载配置文件成功！")
	return server_preferences
}
