package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"wisdom-portal/models"
	"wisdom-portal/schemas"
	"wisdom-portal/wisdom-portal/forms"
	"wisdom-portal/wisdom-portal/result"
)

// 插入监控数据方法
func addMonitor(c *gin.Context) {
	var monitor models.Monitor

	// 绑定表单
	if err := c.ShouldBind(&monitor); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 验证结构
	if err := forms.Validate.Struct(monitor); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, GetValidationError(err.(validator.ValidationErrors))))
		return
	}

	if err := monitor.CreateMonitor(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.NewSuccessResult(result.SuccessCode))
	return
}

// 查询监控数据，返回N个不同tag不同name的监控状态值
// 目前N = 10 OR 20 OR 40 OR 60
func getTcpMonitor(c *gin.Context) {
	var queryMonitor models.TcpQuerySliceMonitor

	// 绑定表单
	if err := c.ShouldBindQuery(&queryMonitor); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 验证结构
	if err := forms.Validate.Struct(queryMonitor); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, GetValidationError(err.(validator.ValidationErrors))))
		return
	}

	// 分页
	if _, num, err := queryMonitor.CountNum(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	} else {
		queryMonitor.Meta = schemas.NewPagination(num)
	}

	// 计算分页
	startNum, endNum := queryMonitor.Meta.PaginationStint(queryMonitor.Page, queryMonitor.PageSize)
	//startNum := (queryMonitor.Page - 1) * schemas.PageSize
	//endNum := queryMonitor.Page * schemas.PageSize

	if err := queryMonitor.QueryMonitor(startNum, endNum); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.NewTcpMonitorQueryResult(result.SuccessCode, queryMonitor, queryMonitor.Meta))
}

// 质量检测
// 监控服务质量检测的方法，返回不同tag不同name的N时间内的监控状态百分比
// N时间 = 1min OR 5min OR 10min OR 20min OR 40min OR 60min
func getTcpMonitorQuality(c *gin.Context) {
	var queryQualityMonitor models.TcpQueryQualitySliceMonitor

	// 绑定表单
	if err := c.ShouldBindQuery(&queryQualityMonitor); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 验证结构
	if err := forms.Validate.Struct(queryQualityMonitor); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, GetValidationError(err.(validator.ValidationErrors))))
		return
	}

	// 分页
	if _, num, err := queryQualityMonitor.CountNum(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	} else {
		queryQualityMonitor.Meta = schemas.NewPagination(num)
	}

	// 计算分页
	startNum, endNum := queryQualityMonitor.Meta.PaginationStint(queryQualityMonitor.Page, queryQualityMonitor.PageSize)
	//startNum := (queryMonitor.Page - 1) * schemas.PageSize
	//endNum := queryMonitor.Page * schemas.PageSize

	if err := queryQualityMonitor.QueryMonitor(startNum, endNum); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.NewTcpMonitorQualityQueryResult(result.SuccessCode, queryQualityMonitor, queryQualityMonitor.Meta))
}
