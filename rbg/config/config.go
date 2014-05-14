package config

import (
	"time"
)

const (
	SAVE_TO_DB_ERROR string = "1001"
	SAVE_BMP_ERROR   string = "1002"
)

type ClientConfig struct {
	// IP地址与端口
	SERVER_IP_PORT string
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
	// 客户端设置名称
	CLIENT_NAME string
	// 日志文件保存路径
	LOG_SAVE_PATH string
}

type ServerConfig struct {
	// IP地址与端口
	SERVER_IP_PORT string
	// 保存BMP的路径
	BMP_SAVE_PATH string
	// 数据库名称
	DATABASE_NAME string
	// 数据库用户名
	DATABASE_USER_NAME string
	// 数据库密码
	DATABASE_PASSWORD string
	// 最大空闲连接数
	DB_MAX_IDLE_CONNS int
	// 最大连接数
	DB_MAX_OPEN_CONNS int
	// 日志文件保存路径
	LOG_SAVE_PATH string
	// 允许连接的IP
	ALLOWS_IP string
}
