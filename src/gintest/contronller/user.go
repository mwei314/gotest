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

type addUserValid struct {
	name string `form:"name" binding:"required"`
	pass string `form:"pass" binding:"required"`
}

// AddUser 新增用户
func AddUser(c *gin.Context) {
	var userValid addUserValid
	if err := c.ShouldBind(userValid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	user := &models.User{
		Name: name,
		Pass: pass,
	}
	err := user.Insert()
	if err != nil {
		c.AbortWithError(501, err)
		return
	}
	c.AbortWithStatusJSON(200, user)
}
