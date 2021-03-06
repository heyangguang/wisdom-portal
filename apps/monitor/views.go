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

// 插入AccessLog数据接口
func addAccessLogMonitor(c *gin.Context) {
	var monitorAccessLog models.MonitorAccessLog

	// 绑定表单
	if err := c.ShouldBind(&monitorAccessLog); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 验证结构
	if err := forms.Validate.Struct(monitorAccessLog); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, GetValidationError(err.(validator.ValidationErrors))))
		return
	}

	if err := monitorAccessLog.CreateMonitor(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.NewSuccessResult(result.SuccessCode))
	return
}

// 插入Client监控数据方法
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

// 插入中间表监控数据
// 插入客户上传状态监控数据 tag 字段区分
func addIntermediateMonitor(c *gin.Context) {
	var monitorIntermediate models.MonitorIntermediate

	// 绑定表单
	if err := c.ShouldBind(&monitorIntermediate); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 验证结构
	if err := forms.Validate.Struct(monitorIntermediate); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, GetValidationError(err.(validator.ValidationErrors))))
		return
	}

	if err := monitorIntermediate.CreateMonitor(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataCreateWrong, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.NewSuccessResult(result.SuccessCode))
	return
}

// 查询中间表监控数据
// 目前N = 10 OR 20 OR 40 OR 60
func queryIntermediateMonitor(c *gin.Context) {
	var queryInMonitor models.QueryIntermediateSliceMonitor

	// 绑定表单
	if err := c.ShouldBindQuery(&queryInMonitor); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 验证结构
	if err := forms.Validate.Struct(queryInMonitor); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, GetValidationError(err.(validator.ValidationErrors))))
		return
	}

	// 分页
	if _, num, err := queryInMonitor.CountNum(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	} else {
		queryInMonitor.Meta = schemas.NewPagination(num)
	}

	// 计算分页
	startNum, endNum := queryInMonitor.Meta.PaginationStint(queryInMonitor.Page, queryInMonitor.PageSize)

	if err := queryInMonitor.QueryMonitor(startNum, endNum); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.NewMonitorQueryInResult(result.SuccessCode, queryInMonitor, queryInMonitor.Meta))
}

// 查询AccessLog数据
func queryAccessLogMonitor(c *gin.Context) {
	var queryAccessMonitor models.QueryAccessLogMonitor

	// 绑定表单
	if err := c.ShouldBindQuery(&queryAccessMonitor); err != nil {
		c.JSON(http.StatusNotAcceptable, result.NewFailResult(result.ParamInvalid, err.Error()))
		return
	}

	// 验证结构
	if err := forms.Validate.Struct(queryAccessMonitor); err != nil {
		c.JSON(http.StatusBadRequest, result.NewSliceFailResult(
			result.ParamInvalid, GetValidationError(err.(validator.ValidationErrors))))
		return
	}

	if err := queryAccessMonitor.QueryMonitor(); err != nil {
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.DataNone, err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.NewMonitorQueryAccessLogResult(result.SuccessCode, queryAccessMonitor))
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
