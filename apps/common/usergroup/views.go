package usergroup

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
	"wisdom-portal/models"
	"wisdom-portal/wisdom-portal/forms"
	"wisdom-portal/wisdom-portal/result"
)

// @Security ApiKeyAuth
// @Summary 添加用户组
// @Description 用于用户组的创建
// @Tags 用户组
// @accept json
// @Produce  json
// @Param data body models.UserGroup true "用户组数据"
// @Success 200 {object} result.SuccessResult "{"code": 10000}"
// @Failure 415 {object} result.FailResult "{"code": 50004}"
// @Failure 400 {object} result.FailResult "{"code": 10001}"
// @Router /api/v1/userGroup [POST]
func addGroup(c *gin.Context) {
	var userGroup models.UserGroup
	// 绑定表单
	if err := c.ShouldBind(&userGroup); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}
	// 验证结构
	if err := forms.Validate.Struct(userGroup); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, GetValidationError(err.(validator.ValidationErrors))))
		return
	}

	// 判断用户名是否存在
	// 判断用户组是否跟用户名相同，不允许相同
	if models.CheckUserName(userGroup.GroupName) || models.CheckUserGroupName(userGroup.GroupName) {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, "group name is already taken"))
		return
	}

	// 防止恶意注入
	userGroup.CreatedAt = time.Now()
	userGroup.UpdatedAt = time.Now()
	if err := userGroup.AddGroup(); err != nil {
		c.JSON(http.StatusUnsupportedMediaType, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.NewSuccessResult(result.SuccessCode))
	return
}
