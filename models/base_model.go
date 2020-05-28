package models

import (
	"grpc_test/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

const (
	PAGE_LIMIT     = 10 //分页每页显示条数，默认值10
)

//开启事务时用这个db(models.db)db.Begin(),db.Rollback(),db.Commit()
var Db *gorm.DB

//时间地区
var Location *time.Location

//连接数据库，创建表
func init() {
	Location, _ = time.LoadLocation("Asia/Shanghai")

	databaseDriver := config.Default("DATABASE_DRIVER", "mysql")
	var err error
	//拼多多数据库
	databaseConnectString := config.Default("DATABASE_CONNECT_STRING", "")
	//打开数据库链接，返回db实例
	Db, err = gorm.Open(databaseDriver, databaseConnectString)
	if err != nil {
		log.Println("错误：", err)
	}

	//开启debug
	//Db.LogMode(true)
}
