package contronller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserList 获取用户列表
func UserList(c *gin.Context) {

	c.String(http.StatusOK, "hello world")
}
