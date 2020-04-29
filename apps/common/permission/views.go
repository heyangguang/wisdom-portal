package permission

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wisdom-portal/models"
)

// 测试
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"context": "perm"})
}

// 添加用户权限
func addPerm(c *gin.Context) {
	casbin := models.CasbinModel{
		RoleName: c.PostForm("role_name"),
		Path:     c.PostForm("path"),
		Method:   c.PostForm("method"),
	}
	isOk := casbin.AddPerm(casbin)
	if isOk {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "保存成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "保存失败"})
	}
	return
}
