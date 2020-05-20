package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
	"wisdom-portal/models"
	wisdomPortal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/forms"
	"wisdom-portal/wisdom-portal/result"
)

// @Security ApiKeyAuth
// @Summary 获取当前用户信息
// @Description 获取当前用户信息
// @Tags 登录管理
// @accept json
// @Produce json
// @Success 200 {object} result.PubCurrentUserResult "{"code": 10000}"
// @Failure 400 {object} result.FailResult "{"code": 20004}"
// @Router /api/v1/pub/current/user [GET]
func getCurrentUser(c *gin.Context) {
	if username, isExist := c.Get("username"); isExist {
		var pubCurrentUser models.PubCurrentUser
		if pubCurrentUserObj, isDbExist := pubCurrentUser.GetPubCurrentUser(username); isDbExist {
			c.JSON(http.StatusOK, result.NewPubCurrentUserResult(result.SuccessCode, *pubCurrentUserObj))
			return
		}
	}
	c.JSON(http.StatusBadRequest, result.NewFailResult(result.UserNotExist, ""))
	return
}

// @Summary 注册用户
// @Description 用于用户的注册
// @Tags 用户注册
// @accept json
// @Produce json
// @Param data body models.User true "用户注册数据"
// @Success 200 {object} result.RegisterUserResult "{"code": 10000}"
// @Failure 415 {object} result.FailResult "{"code": 50004}"
// @Failure 415 {object} result.FailResult "{"code": 50003}"
// @Failure 400 {object} result.FailResult "{"code": 10001}"
// @Router /api/v1/register [POST]
func register(c *gin.Context) {
	var user models.User
	// 绑定表单
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}
	// 验证结构
	if err := forms.Validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, GetValidationError(err.(validator.ValidationErrors))))
		return
	}

	// 判断用户名是否存在
	// 判断用户组是否跟用户名相同，不允许相同
	if models.CheckUserName(user.UserName) || models.CheckUserGroupName(user.UserName) {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, "username is already taken"))
		return
	}

	// 防止恶意数据请求
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Status = true
	user.PassWord = wisdomPortal.String2md5(user.PassWord)
	// 双因子认证TOTP
	var googleAuth *models.GoogleAuth
	googleAuth = models.NewGoogleAuth()
	user.Secret = googleAuth.GetSecret()
	googleAuth.GetQrCode(user.UserName, user.Secret)
	if err := user.AddUser(user); err != nil {
		c.JSON(http.StatusUnsupportedMediaType, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}
	if err := models.AddDefaultPerm(user.UserName, "/api/v1/pub/current/user", "*"); err != nil {
		c.JSON(http.StatusUnsupportedMediaType, result.NewFailResult(result.DataAlreadyExisted, err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.NewRegisterUserResult(result.SuccessCode, *googleAuth))
	return
}
