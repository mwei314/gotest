package models

import (
	"gintest/utils"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// GormDB MySQL连接
var GormDB *gorm.DB

// Init 初始化
func Init(conf *utils.Config) {
	gormDB, err := gorm.Open("mysql", conf.Mysql)
	if err != nil {
		panic(err)
	}
	// 设置最大连接数
	gormDB.DB().SetMaxOpenConns(10)
	// 设置最大空闲连接数
	gormDB.DB().SetMaxIdleConns(10)
	// 设置每个链接的过期时间
	gormDB.DB().SetConnMaxLifetime(time.Second * 5)
	GormDB = gormDB

	// 初始化数据库
	GormDB.AutoMigrate(&User{})
}
