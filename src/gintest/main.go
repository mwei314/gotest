package main

import (
	"gintest/libs/config"
	"gintest/libs/logger"
	"gintest/libs/redis"
	"gintest/models"
	"gintest/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Init()
	redis.Init(&config.Redis)
	models.Init(config.Mysql)
	logger.Init(config.Logger.Path)
	// 初始化router
	app := gin.New()
	router.Init(app)
	//启动服务
	app.Run(":8080")
	logger.Info("server start!")
}
