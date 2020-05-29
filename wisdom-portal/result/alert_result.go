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

func NewAlertQueryResult(code int, data models.QuerySliceAlert, meta *schemas.Pagination) *AlertQueryResult {
	return &AlertQueryResult{
		Code: code,
		Msg:  ResultText(code),
		Meta: *meta,
		Data: data.Data,
	}
}
