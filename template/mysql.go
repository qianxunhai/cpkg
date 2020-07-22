package template

var Mysql =`package gmysql

import (
	"{{.Name}}/pkg/logging"
	"{{.Name}}/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	logging.AppLogger.Info("setup mysql...")
	//读取配置文件
	conf := "database"

	setting.MapTo(conf, DatabaseSetting)

	var err error
	DB, err = gorm.Open(DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.Name))

	if err != nil {
		logging.AppLogger.Error("models.Setup err:"+err.Error())
	}
	//更改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return DatabaseSetting.TablePrefix + defaultTableName
	}
	// 全局禁用表名复数
	DB.SingularTable(true)
	//设置数据库连接池
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	logging.AppLogger.Info("setup mysql complete")
}

//关闭数据库
func CloseDB() {
	defer DB.Close()
}
`
