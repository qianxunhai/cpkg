package template

var MainFunc =`package main

import (
    "{{.Name}}/pkg/gmysql"
    "{{.Name}}/pkg/gredis"
    "{{.Name}}/pkg/logging"
    "{{.Name}}/pkg/setting"
    "{{.Name}}/routers"
)


//系统初始化
func init() {
	//全局设置
	setting.Setup()
	//日志配置
	logging.Setup()
	//mysql配置
    gmysql.Setup()
	//redis配置
	gredis.Setup()

}

func main() {
	r:=routers.InitRoute()
	logging.AppLogger.Info("start "+setting.ServerSetting.HttpHost+":"+setting.ServerSetting.HttpPort)
	r.Run(setting.ServerSetting.HttpHost+":"+setting.ServerSetting.HttpPort)
}`
