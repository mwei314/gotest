package controllers

import (
	"gintest/libs/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 首页控制器
func Index(c *gin.Context) {
	logger.Info("index page")
	c.String(http.StatusOK, "hello world")
}

// returnJSON 通用的返回json结果方法
func returnJSON(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
