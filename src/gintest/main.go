package main

import (
	"gintest/models"
	"gintest/redis"
	"gintest/utils"
	"gintest/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config := utils.InitConfig()
	redis.Init(config)
	models.Init(config)
	// 初始化router
	app := gin.New()
	router.Init(app)
	//启动服务
	app.Run(":8080")
}
