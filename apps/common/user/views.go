package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wisdom-portal/models"
	wisdomPortal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/result"
)

// 测试
func test(c *gin.Context) {
	var user1 models.User
	models.DB.Preload("UserGroups").First(&user1, "id = ?", 1)
	c.JSON(http.StatusOK, gin.H{"data": user1})
}

// @Summary 注册用户
// @Description 用于用户的注册
// @Tags 用户注册
// @accept json
// @Produce  json
// @Param data body models.SwaggerUser true "用户注册数据"
// @Success 200 {object} gin.H "{"code": 10000, "msg": "成功", "data": models.GoogleAuth}"
// @Failure 415 {object} gin.H "{"code": 50004, "msg": "数据创建错误", "err": ""}"
// @Failure 400 {object} gin.H "{"code": 10001, "msg": "参数无效", "err": ""}"
// @Router /api/v1/register [POST]
func register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": result.ParamInvalid,
			"msg":  result.ResultText(result.ParamInvalid),
			"err":  err.Error(),
		})
		return
	}
	// 防止恶意数据请求
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Status = false
	user.PassWord = wisdomPortal.String2md5(user.PassWord)
	// 双因子认证TOTP
	var googleAuth *models.GoogleAuth
	googleAuth = models.NewGoogleAuth()
	user.Secret = googleAuth.GetSecret()
	googleAuth.GetQrCode(user.UserName, user.Secret)
	if err := user.AddUser(user); err != nil {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"code": result.DataCreateWrong,
			"msg":  result.ResultText(result.DataCreateWrong),
			"err":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": result.SuccessCode,
		"msg":  result.ResultText(result.SuccessCode),
		"data": googleAuth,
	})
	return
}
