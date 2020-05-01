package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wisdom-portal/models"
	"wisdom-portal/wisdom-portal/result"
)

// @Summary 用户登录
// @Description 用于用户登录
// @Tags 用户登录
// @accept json
// @Produce  json
// @Param data body models.UserLogin true "数据"
// @Success 200 {object} gin.H "{"code": 10000, "msg": "成功", "data": "{"token": ""}"}"
// @Failure 400 {object} gin.H "{"code": 10001, "msg": "参数无效", "err": ""}"
// @Failure 401 {object} gin.H "{"code": 20002, "msg": "账号不存在或密码动态码错误"}"
// @Router /api/v1/login [POST]
func login(c *gin.Context) {
	// 用户发送用户名和密码过来
	var userLogin models.UserLogin
	err := c.ShouldBind(&userLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": result.ParamInvalid,
			"msg":  result.ResultText(result.ParamInvalid),
			"err":  err.Error(),
		})
		return
	}
	// 校验用户名和密码是否正确
	if userLogin.CheckUserLogin() && userLogin.CheckUserOtpCode() {
		// 生成Token
		tokenString := models.GenToken(userLogin.UserName)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": result.SuccessCode,
			"msg":  result.ResultText(result.SuccessCode),
			"data": gin.H{"token": tokenString},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": result.UserLoginError,
		"msg":  result.ResultText(result.UserLoginError),
	})
	return
}
