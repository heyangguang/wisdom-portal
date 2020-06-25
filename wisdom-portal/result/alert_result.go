package result

import (
	"wisdom-portal/models"
	"wisdom-portal/schemas"
)

// 监控
type AlertQueryResult struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Meta schemas.Pagination  `json:"meta"`
	Data []models.QueryAlert `json:"data"`
}

// 监控等级条数
type AlertQueryCountResult struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data models.QueryCountAlert `json:"data"`
}

func NewAlertQueryResult(code int, data models.QuerySliceAlert, meta *schemas.Pagination) *AlertQueryResult {
	return &AlertQueryResult{
		Code: code,
		Msg:  ResultText(code),
		Meta: *meta,
		Data: data.Data,
	}
}

func NewAlertQueryCountResult(code int, data models.QueryCountAlert) *AlertQueryCountResult {
	return &AlertQueryCountResult{
		Code: code,
		Msg:  ResultText(code),
		Data: data,
	}
}
