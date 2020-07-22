package template

var Setting =`package setting

import (
	"go.uber.org/zap"
	"log"
	"time"
	"github.com/go-ini/ini"
)

//服务相关
type Server struct {
	ProjectName  string
	RunMode      string
	HttpPort     string
	HttpHost     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("setting.Setup, fail to parse 'conf/app.ini' ", zap.Error(err))
	}

	MapTo("server", ServerSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func MapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatal("Cfg.MapTo Setting err' ", zap.Error(err))
	}
}
`
