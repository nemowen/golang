package config

import (
	"time"
)

type ClientConfig struct {
	// IP地址与端口
	ADDR_PORT string
	// 监控文件
	FLAG_FILE_PATH string
	// 采集数据文件
	NOTE_FILE_PATH string
	// 开始工作状态
	START_WORK_FLAG string
	// 读取flag.ini文件速度， 称为单位
	FLAG_STATE_SPEED time.Duration
	// 当连接失败时，多少秒后重新连接服务器
	RECONNECT_SERVER_TIME time.Duration
}

type ServerConfig struct {
	// IP地址与端口
	ADDR_PORT string
	// 保存BMP的路径
	BMP_SAVE_PATH string
}