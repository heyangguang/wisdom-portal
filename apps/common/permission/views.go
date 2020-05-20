package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
	"wisdom-portal/models"
	"wisdom-portal/wisdom-portal/forms"
	"wisdom-portal/wisdom-portal/result"
)

// 测试
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"context": "perm"})
}

// @Security ApiKeyAuth
// @Summary 添加权限模板
// @Description 用于创建权限模板
// @Tags 权限
// @accept json
// @Produce  json
// @Param data body models.Rule true "数据"
// @Success 201 {object} result.RegisterUserResult "{"code": 10000}"
// @Failure 406 {object} result.FailResult "{"code": 10001}"
// @Failure 400 {object} result.SliceFailResult "{"code": 10001}"
// @Failure 500 {object} result.FailResult "{"code": 50004}"
// @Router /api/v1/perm [POST]
func addPerm(c *gin.Context) {
	var rule models.Rule

	// 绑定表单
	if err := c.ShouldBind(&rule); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}
	// 验证结构
	if err := forms.Validate.Struct(rule); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, forms.BaseFormValidationError(err.(validator.ValidationErrors))))
		return
	}

	// 防止恶意数据请求
	rule.CreatedAt = time.Now()
	rule.UpdatedAt = time.Now()

	err := rule.AddPerm(rule)
	if err == nil {
		c.JSON(http.StatusCreated, result.NewSuccessResult(result.SuccessCode))
		return
	}
	c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataCreateWrong, err.Error()))
	return
}

// @Security ApiKeyAuth
// @Summary 用户关联权限模板
// @Description 用于用户关联权限
// @Tags 权限
// @accept json
// @Produce  json
// @Param rule_id body models.AddPermUser true "权限模板id"
// @Param uid path string true "用户id"
// @Success 201 {object} result.SuccessResult "{"code": 10000}"
// @Failure 412 {object} result.FailResult "{"code": 10004}"
// @Failure 406 {object} result.FailResult "{"code": 10001}"
// @Failure 400 {object} result.SliceFailResult "{"code": 10001}"
// @Failure 500 {object} result.FailResult "{"code": 50004}"
// @Router /api/v1/perm/user/{uid} [POST]
func addPermUser(c *gin.Context) {
	var addPermUser models.AddPermUser
	uid := c.Param("uid")
	if uid == "" {
		c.JSON(http.StatusPreconditionFailed, result.NewFailResult(result.ParamNotComplete, ""))
		return
	}

	// 绑定表单
	if err := c.ShouldBind(&addPermUser); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}
	// 验证结构
	if err := forms.Validate.Struct(addPermUser); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, forms.BaseFormValidationError(err.(validator.ValidationErrors))))
		return
	}

	if err := addPermUser.AddPermUser(uid); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, result.NewSuccessResult(result.SuccessCode))
	return
}

// @Security ApiKeyAuth
// @Summary 用户组关联权限模板
// @Description 用于用户组关联权限
// @Tags 权限
// @accept json
// @Produce  json
// @Param rule_id body models.AddPermUserGroup true "权限模板id"
// @Param gid path string true "用户组id"
// @Success 201 {object} result.SuccessResult "{"code": 10000}"
// @Failure 412 {object} result.FailResult "{"code": 10004}"
// @Failure 406 {object} result.FailResult "{"code": 10001}"
// @Failure 400 {object} result.SliceFailResult "{"code": 10001}"
// @Failure 500 {object} result.FailResult "{"code": 50004}"
// @Router /api/v1/perm/userGroup/{gid} [POST]
func addPermUserGroup(c *gin.Context) {
	var addPermUserGroup models.AddPermUserGroup
	gid := c.Param("gid")
	if gid == "" {
		c.JSON(http.StatusPreconditionFailed, result.NewFailResult(result.ParamNotComplete, ""))
	}

	// 绑定表单
	if err := c.ShouldBind(&addPermUserGroup); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}
	// 验证结构
	if err := forms.Validate.Struct(addPermUserGroup); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, forms.BaseFormValidationError(err.(validator.ValidationErrors))))
		return
	}

	if err := addPermUserGroup.AddPermUserGroup(gid); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, result.NewSuccessResult(result.SuccessCode))
	return
}
