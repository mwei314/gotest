package router

import (
	"gintest/controllers"

	"github.com/gin-gonic/gin"
)

// Init 路由初始化
func Init(router *gin.Engine) {
	router.GET("/", controllers.Index)
	router.GET("/redis/set", controllers.RedisSet)
	router.GET("/redis/get", controllers.RedisGet)

	// 用户相关
	router.POST("/user", controllers.AddUser)
}
