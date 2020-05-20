package permission

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wisdom-portal/models"
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
// @Success 200 {object} result.SuccessResult "{"code": 10000}"
// @Failure 415 {object} result.FailResult "{"code": 50004}"
// @Failure 400 {object} result.FailResult "{"code": 10001}"
// @Router /api/v1/perm [POST]
func addPerm(c *gin.Context) {
	var rule models.Rule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 防止恶意数据请求
	rule.CreatedAt = time.Now()
	rule.UpdatedAt = time.Now()

	err := rule.AddPerm(rule)
	if err == nil {
		c.JSON(http.StatusOK, result.NewSuccessResult(result.SuccessCode))
		return
	}
	c.JSON(http.StatusUnsupportedMediaType, result.NewFailResult(result.DataCreateWrong, err.Error()))
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
// @Success 200 {object} result.SuccessResult "{"code": 10000}"
// @Failure 415 {object} result.FailResult "{"code": 50004}"
// @Failure 400 {object} result.FailResult "{"code": 10001}"
// @Failure 400 {object} result.FailResult "{"code": 10004}"
// @Router /api/v1/perm/user/{uid} [POST]
func addPermUser(c *gin.Context) {
	var addPermUser models.AddPermUser
	uid := c.Param("uid")
	if uid == "" {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamNotComplete, ""))
		return
	}
	if err := c.ShouldBindJSON(&addPermUser); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	if err := addPermUser.AddPermUser(uid); err != nil {
		c.JSON(http.StatusUnsupportedMediaType, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.NewSuccessResult(result.SuccessCode))
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
// @Success 200 {object} result.SuccessResult "{"code": 10000}"
// @Failure 415 {object} result.FailResult "{"code": 50004}"
// @Failure 400 {object} result.FailResult "{"code": 10001}"
// @Failure 400 {object} result.FailResult "{"code": 10004}"
// @Router /api/v1/perm/userGroup/{gid} [POST]
func addPermUserGroup(c *gin.Context) {
	var addPermUserGroup models.AddPermUserGroup
	gid := c.Param("gid")
	if gid == "" {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamNotComplete, ""))
	}
	if err := c.ShouldBindJSON(&addPermUserGroup); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	if err := addPermUserGroup.AddPermUserGroup(gid); err != nil {
		c.JSON(http.StatusUnsupportedMediaType, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.NewSuccessResult(result.SuccessCode))
	return
}
