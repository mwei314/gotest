package router

import (
	"gintest/contronller"

	"github.com/gin-gonic/gin"
)

// Init 路由初始化
func Init(router *gin.Engine) {
	router.GET("/", contronller.Index)
	router.GET("/redis/set", contronller.RedisSet)
	router.GET("/redis/get", contronller.RedisGet)
}
