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
		returnJSON(c, 1, "参数错误", nil)
		return
	}
	name := c.PostForm("name")
	checkUser := models.FindUserByName(name)
	if checkUser.ID != 0 {
		returnJSON(c, 1, "姓名重复", nil)
		return
	}
	pass := c.PostForm("pass")
	user := &models.User{
		Name: name,
		Pass: pass,
	}
	err := user.Insert()
	if err != nil {
		// TODO 处理error
		returnJSON(c, 1, "新增失败", nil)
		return
	}
	returnJSON(c, 0, "", user)
}
