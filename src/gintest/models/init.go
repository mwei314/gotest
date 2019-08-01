package models

import (
	"database/sql"
	"gintest/utils"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB 数据库操作结构体
var DB *sql.DB

// Init 初始化
func Init(conf *utils.Config) {
	db, err := sql.Open("mysql", conf.Mysql)
	if err != nil {
		panic(err)
	}
	// 设置最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(10)
	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(time.Second * 5)
	DB = db
}
