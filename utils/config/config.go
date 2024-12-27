package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	Dbname     string

	QQEmail   string
	QQGenCode string
)

func init() {
	// 这里是相对路径
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}

	// 加载配置文件，读取其中的值设置全局变量
	LoadServer(file)
	LoadData(file)
	LoadEmail(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").String()
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").String()
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassword = file.Section("database").Key("DbPassword").String()
	Dbname = file.Section("database").Key("DbName").String()
}

func LoadEmail(file *ini.File) {
	QQEmail = file.Section("email").Key("QQEmail").String()
	QQGenCode = file.Section("email").Key("QQGenCode").String()
}
