package controllers

import (
	"gintest/libs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RedisGet 测试redis的get操作
func RedisGet(c *gin.Context) {
	info, err := libs.RedisDo("GET", "test")
	if err != nil {
		c.String(http.StatusExpectationFailed, "error")
		return
	}
	if info == nil {
		c.String(http.StatusExpectationFailed, "none")
		return
	}
	c.String(http.StatusOK, "%s", info)
}

// RedisSet 测试redis的set操作
func RedisSet(c *gin.Context) {
	_, err := libs.RedisDo("SET", "test", "test", "EX", 1000)
	if err != nil {
		c.String(http.StatusExpectationFailed, "error")
		return
	}
	c.String(http.StatusOK, "done")
}
