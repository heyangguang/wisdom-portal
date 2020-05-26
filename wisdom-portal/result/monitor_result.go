package result

import (
	"wisdom-portal/models"
	"wisdom-portal/schemas"
)

type MonitorQueryResult struct {
	Code int                                `json:"code"`
	Msg  string                             `json:"msg"`
	Meta schemas.Pagination                 `json:"meta"`
	Data []map[string][]models.QueryMonitor `json:"data"`
}

func NewMonitorQueryResult(code int, data models.QuerySliceMonitor, meta *schemas.Pagination) *MonitorQueryResult {
	return &MonitorQueryResult{
		Code: code,
		Msg:  ResultText(code),
		Meta: *meta,
		Data: data.Data,
	}
}
