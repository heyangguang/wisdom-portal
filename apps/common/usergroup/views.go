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
// @Success 200 {object} gin.H "{"code": 10000, "msg": "成功"}"
// @Failure 415 {object} gin.H "{"code": 50004, "msg": "数据创建错误", "err": ""}"
// @Failure 400 {object} gin.H "{"code": 10001, "msg": "参数无效", "err": ""}"
// @Router /api/v1/userGroup [POST]
func addGroup(c *gin.Context) {
	var userGroup models.UserGroup
	if err := c.ShouldBindJSON(&userGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": result.ParamInvalid,
			"msg":  result.ResultText(result.ParamInvalid),
			"err":  err.Error(),
		})
		return
	}
	fmt.Println(userGroup.Users)
	// 防止恶意注入
	userGroup.CreatedAt = time.Now()
	userGroup.UpdatedAt = time.Now()
	if err := userGroup.AddGroup(); err != nil {
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
	})
	return
}
