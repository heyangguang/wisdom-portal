package usergroup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wisdom-portal/models"
)

// 添加用户组
func addGroup(c *gin.Context) {
	var userGroup models.UserGroup
	if err := c.ShouldBindJSON(&userGroup); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "表单填写错误，请检查后重新提交",
			"err":  err.Error(),
		})
		return
	}
	fmt.Println(userGroup.Users)
	// 防止恶意注入
	userGroup.CreatedAt = time.Now()
	userGroup.UpdatedAt = time.Now()
	if err := userGroup.AddGroup(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建用户组失败",
			"err":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "创建用户组成功"})
	return
}
