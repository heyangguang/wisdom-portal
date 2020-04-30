package permission

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wisdom-portal/models"
)

// 测试
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"context": "perm"})
}

// 添加权限模板
func addPerm(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "表单填写错误，请检查后重新提交",
			"err":  err.Error(),
		})
		return
	}

	// 防止恶意数据请求
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	err := role.AddPerm(role)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "保存成功"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "创建模板权限失败", "err": err.Error()})
	return
}

// 用户对应权限模板
func addPermUser(c *gin.Context) {
	var addPermUser models.AddPermUser
	uid := c.Param("uid")
	if uid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "选择用户失败",
		})
		return
	}
	if err := c.ShouldBindJSON(&addPermUser); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "表单填写错误，请检查后重新提交",
			"err":  err.Error(),
		})
		return
	}

	if err := addPermUser.AddPermUser(uid); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "用户权限添加错误", "err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "保存成功"})
	return
}
