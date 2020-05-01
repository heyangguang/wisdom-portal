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
// @Param data body models.SwaggerRole true "数据"
// @Success 200 {object} gin.H "{"code": 10000, "msg": "成功", "data": ""}"
// @Failure 415 {object} gin.H "{"code": 50004, "msg": "数据创建错误", "err": ""}"
// @Failure 400 {object} gin.H "{"code": 10001, "msg": "参数无效"}"
// @Router /api/v1/perm [POST]
func addPerm(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": result.ParamInvalid,
			"msg":  result.ResultText(result.ParamInvalid),
			"err":  err.Error(),
		})
		return
	}

	// 防止恶意数据请求
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	err := role.AddPerm(role)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": result.SuccessCode, "msg": result.ResultText(result.SuccessCode)})
		return
	}
	c.JSON(http.StatusUnsupportedMediaType, gin.H{
		"code": result.DataCreateWrong,
		"msg":  result.ResultText(result.DataCreateWrong),
		"err":  err.Error(),
	})
	return
}

// @Security ApiKeyAuth
// @Summary 用户关联权限模板
// @Description 用于用户关联权限
// @Tags 权限
// @accept json
// @Produce  json
// @Param role_id body models.AddPermUser true "权限模板id"
// @Param uid path string true "用户id"
// @Success 200 {object} gin.H "{"code": 10000, "msg": "成功", "data": ""}"
// @Failure 415 {object} gin.H "{"code": 50004, "msg": "数据创建错误", "err": ""}"
// @Failure 400 {object} gin.H "{"code": 10001, "msg": "参数无效", "err": ""}"
// @Failure 400 {object} gin.H "{"code": 10004, "msg": "参数缺失"}"
// @Router /api/v1/perm/user/{uid} [POST]
func addPermUser(c *gin.Context) {
	var addPermUser models.AddPermUser
	uid := c.Param("uid")
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": result.ParamNotComplete,
			"msg":  result.ResultText(result.ParamNotComplete),
		})
		return
	}
	if err := c.ShouldBindJSON(&addPermUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": result.ParamInvalid,
			"msg":  result.ResultText(result.ParamInvalid),
			"err":  err.Error(),
		})
		return
	}

	if err := addPermUser.AddPermUser(uid); err != nil {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"code": result.DataCreateWrong,
			"msg":  result.ResultText(result.DataCreateWrong),
			"err":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.SuccessCode, "msg": result.ResultText(result.SuccessCode)})
	return
}

// @Security ApiKeyAuth
// @Summary 用户组关联权限模板
// @Description 用于用户组关联权限
// @Tags 权限
// @accept json
// @Produce  json
// @Param role_id body models.AddPermUserGroup true "权限模板id"
// @Param gid path string true "用户组id"
// @Success 200 {object} gin.H "{"code": 10000, "msg": "成功", "data": ""}"
// @Failure 415 {object} gin.H "{"code": 50004, "msg": "数据创建错误", "err": ""}"
// @Failure 400 {object} gin.H "{"code": 10001, "msg": "参数无效", "err": ""}"
// @Failure 400 {object} gin.H "{"code": 10004, "msg": "参数缺失"}"
// @Router /api/v1/perm/userGroup/{gid} [POST]
func addPermUserGroup(c *gin.Context) {
	var addPermUserGroup models.AddPermUserGroup
	gid := c.Param("gid")
	if gid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": result.ParamNotComplete,
			"msg":  result.ResultText(result.ParamNotComplete),
		})
	}
	if err := c.ShouldBindJSON(&addPermUserGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": result.ParamInvalid,
			"msg":  result.ResultText(result.ParamInvalid),
			"err":  err.Error(),
		})
		return
	}

	if err := addPermUserGroup.AddPermUserGroup(gid); err != nil {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"code": result.DataCreateWrong,
			"msg":  result.ResultText(result.DataCreateWrong),
			"err":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": result.SuccessCode, "msg": result.ResultText(result.SuccessCode)})
	return
}
