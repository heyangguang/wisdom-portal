package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"wisdom-portal/apps/common/user"
	"wisdom-portal/models"
	"wisdom-portal/wisdom-portal/forms"
	"wisdom-portal/wisdom-portal/result"
)

// @Summary 用户登录
// @Description 用于用户登录
// @Tags 登录管理
// @accept json
// @Produce  json
// @Param data body models.UserLogin true "数据"
// @Success 200 {object} result.LoginResult "{"code": 10000}"
// @Failure 400 {object} result.FailResult "{"code": 10001}"
// @Failure 401 {object} result.FailResult "{"code": 20002}"
// @Router /api/v1/login [POST]
func login(c *gin.Context) {
	// 用户发送用户名和密码过来
	var userLogin models.UserLogin
	// 绑定表单
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}
	// 验证结构
	if err := forms.Validate.Struct(userLogin); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, user.GetValidationError(err.(validator.ValidationErrors))))
		return
	}
	// 校验用户名和密码是否正确
	if userLogin.CheckUserLogin() && userLogin.CheckUserOtpCode() {
		// 生成Token
		tokenString := models.GenToken(userLogin.UserName)
		c.JSON(http.StatusOK, result.NewLoginResult(result.SuccessCode, *result.NewLoginTokenResult(tokenString)))
		return
	}
	// 登录失败
	c.JSON(http.StatusUnauthorized, result.NewFailResult(result.UserLoginError, ""))
	return
}
