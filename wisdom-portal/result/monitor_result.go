package result

import "wisdom-portal/models"

type MonitorQueryResult struct {
	Code int                                `json:"code"`
	Msg  string                             `json:"msg"`
	Data []map[string][]models.QueryMonitor `json:"data"`
}

func NewMonitorQueryResult(code int, data models.QuerySliceMonitor) *MonitorQueryResult {
	return &MonitorQueryResult{
		Code: code,
		Msg:  ResultText(code),
		Data: data.Data,
	}
}
