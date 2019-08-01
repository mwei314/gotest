package contronller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 首页控制器
func Index(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
