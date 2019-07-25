package main

import (
	"gintest/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化router
	router := gin.New()
	redis.Init()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	router.GET("/redis/set", func(c *gin.Context) {
		_, err := redis.Do("SET", "test", "test", "EX", 1000)
		if err != nil {
			c.String(http.StatusExpectationFailed, "error")
			return
		}
		c.String(http.StatusOK, "done")
	})
	router.GET("/redis/get", func(c *gin.Context) {
		info, err := redis.Do("GET", "test")
		if err != nil {
			c.String(http.StatusExpectationFailed, "error")
			return
		}
		if info == nil {
			c.String(http.StatusExpectationFailed, "none")
			return
		}
		c.String(http.StatusOK, "%s", info)
	})
	//启动服务
	router.Run(":8080")
}
