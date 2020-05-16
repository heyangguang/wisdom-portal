package usergroup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wisdom-portal/models"
	"wisdom-portal/wisdom-portal/result"
)

// @Security ApiKeyAuth
// @Summary 添加用户组
// @Description 用于用户组的创建
// @Tags 用户组
// @accept json
// @Produce  json
// @Param data body models.SwaggerUserGroup true "用户组数据"
// @Success 200 {object} result.SuccessResult "{"code": 10000}"
// @Failure 415 {object} result.FailResult "{"code": 50004}"
// @Failure 400 {object} result.FailResult "{"code": 10001}"
// @Router /api/v1/userGroup [POST]
func addGroup(c *gin.Context) {
	var userGroup models.UserGroup
	if err := c.ShouldBindJSON(&userGroup); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}
	fmt.Println(userGroup.Users)
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
