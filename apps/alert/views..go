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

// 查询告警状态
func queryAlert(c *gin.Context) {
	var querySliceAlert models.QuerySliceAlert

	// 绑定表单
	if err := c.ShouldBindQuery(&querySliceAlert); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 验证结构
	if err := forms.Validate.Struct(querySliceAlert); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, forms.BaseFormValidationError(err.(validator.ValidationErrors))))
		return
	}

	// 分页
	if num, err := querySliceAlert.CountNum(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	} else {
		logger.Debug(fmt.Sprintf("count_num: %d", num))
		querySliceAlert.Meta = schemas.NewPagination(num)
	}

	// 计算分页
	startNum, endNum := querySliceAlert.Meta.PaginationStint(querySliceAlert.Page, querySliceAlert.PageSize)
	logger.Debug(fmt.Sprintf("startNum: %d", startNum))
	logger.Debug(fmt.Sprintf("endNum: %d", endNum))

	if err := querySliceAlert.QueryAlert(startNum, endNum); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.NewAlertQueryResult(result.SuccessCode, querySliceAlert, querySliceAlert.Meta))
}

// 查询告警条数
func queryCountAlert(c *gin.Context) {
	var queryCountAlert models.QueryCountAlert
	if err := queryCountAlert.LevelCount(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	}
	fmt.Println(queryCountAlert.LevelOne)
	c.JSON(http.StatusOK, result.NewAlertQueryCountResult(result.SuccessCode, queryCountAlert))
}
