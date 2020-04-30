package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wisdom-portal/models"
	wisdomPortal "wisdom-portal/wisdom-portal"
)

// @Summary 测试接口
// @Description get data
// @Accept  json
// @Produce json
// @Success 200 {object} models.User "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/user [get]
// 测试
func test(c *gin.Context) {
	var user1 models.User
	models.DB.Preload("UserGroups").First(&user1, "id = ?", 1)
	c.JSON(http.StatusOK, gin.H{"data": user1})
}

// @Summary 注册用户
func register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "表单填写错误，请检查后重新提交",
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
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建用户失败，请检查后重新提交",
			"err":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "创建用户成功",
		"data": googleAuth,
	})
	return
}
