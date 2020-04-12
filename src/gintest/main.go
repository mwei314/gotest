package main

import (
	"gintest/libs"
	"gintest/models"
	"gintest/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config := libs.InitConfig()
	libs.RedisInit(config)
	models.Init(config)
	// 初始化router
	app := gin.New()
	router.Init(app)
	//启动服务
	app.Run(":8080")
}
