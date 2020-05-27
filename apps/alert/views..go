package alert

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"wisdom-portal/models"
	"wisdom-portal/schemas"
	"wisdom-portal/wisdom-portal/forms"
	"wisdom-portal/wisdom-portal/logger"
	"wisdom-portal/wisdom-portal/result"
)

// 获取prometheus alertManager中的告警方法
func addAlert(c *gin.Context) {
	var alertManager schemas.AlertManagerWebHook
	var alert models.Alert

	// 绑定表单
	if err := c.ShouldBind(&alertManager); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 验证结构
	if err := forms.Validate.Struct(alertManager); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, forms.BaseFormValidationError(err.(validator.ValidationErrors))))
		return
	}

	logger.Debug(fmt.Sprintf("%v", alertManager))

	// 执行插入
	if err := alert.CreateAlert(&alertManager); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.NewSuccessResult(result.SuccessCode))
}

// 修改告警状态
func updateAlert(c *gin.Context) {
	var alert models.Alert
	id := c.Param("id")
	// 验证id是否有值
	if id == "" {
		c.JSON(http.StatusPreconditionFailed, result.NewFailResult(result.ParamNotComplete, ""))
		return
	}

	// 判断数据是否存在
	if err := alert.QueryId(id); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.DataNone, err.Error()))
		return
	}

	// 执行更新
	if err := alert.UpdateAlert(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.NewSuccessResult(result.SuccessCode))
}
