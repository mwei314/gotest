package contronller

import (
	"net/http"

	"gintest/models"

	"github.com/gin-gonic/gin"
)

// UserList 获取用户列表
func UserList(c *gin.Context) {

	c.String(http.StatusOK, "hello world")
}

// AddUser 新增用户
func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	user := &models.User{
		Name: name,
		Pass: pass,
	}
	err := user.Insert()
	if err != nil {
		c.AbortWithError(501, err)
	}
	c.AbortWithStatusJSON(200, user)
}
