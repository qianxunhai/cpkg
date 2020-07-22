package template

var Log_log =`package logging

import (
	"{{.Name}}/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Log配置结构体
type Log struct {
	Gin         string
	App         string
	Gateway     string
	ServiceName string
}

var LogSetting = &Log{}
var AppLogger *zap.Logger
var GatewayLogger *zap.Logger

//定制日志
func Setup() {
	setting.MapTo("log", LogSetting)
	//记录Gin日志
	//f, _ := os.Create(LogSetting.Gin)
	// Use the following code if you need to write the logs to file and console at the same time.
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//定制日志
	AppLogger = NewLogger(LogSetting.App, zapcore.InfoLevel, 128, 30, 7, true, LogSetting.ServiceName)
	//gateway
	//GatewayLogger = NewLogger(LogSetting.Gateway, zapcore.InfoLevel, 128, 30, 7, true, LogSetting.ServiceName)
}
`
