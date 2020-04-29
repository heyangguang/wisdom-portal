package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wisdom-portal/models"
)

func login(c *gin.Context) {
	// 用户发送用户名和密码过来
	var userLogin models.UserLogin
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "无效的参数, err:" + err.Error(),
		})
		return
	}
	// 校验用户名和密码是否正确
	if userLogin.CheckUserLogin() && userLogin.CheckUserOtpCode() {
		// 生成Token
		tokenString := models.GenToken(userLogin.UserName)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  "账号或密码错误，请重新输入",
	})
	return
}
