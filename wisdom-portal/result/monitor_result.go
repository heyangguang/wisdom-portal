package result

import (
	"wisdom-portal/models"
	"wisdom-portal/schemas"
)

// 服务状态
type TcpMonitorQueryResult struct {
	Code int                                   `json:"code"`
	Msg  string                                `json:"msg"`
	Meta schemas.Pagination                    `json:"meta"`
	Data []map[string][]models.TcpQueryMonitor `json:"data"`
}

// 服务质量
type TcpMonitorQualityQueryResult struct {
	Code int                             `json:"code"`
	Msg  string                          `json:"msg"`
	Meta schemas.Pagination              `json:"meta"`
	Data []models.TcpQueryQualityMonitor `json:"data"`
}

// 中间表
type MonitorQueryInResult struct {
	Code int                                       `json:"code"`
	Msg  string                                    `json:"msg"`
	Meta schemas.Pagination                        `json:"meta"`
	Data []map[string][]models.MonitorIntermediate `json:"data"`
}

// 服务状态
func NewTcpMonitorQueryResult(code int, data models.TcpQuerySliceMonitor, meta *schemas.Pagination) *TcpMonitorQueryResult {
	return &TcpMonitorQueryResult{
		Code: code,
		Msg:  ResultText(code),
		Meta: *meta,
		Data: data.Data,
	}
}

// 服务质量
func NewTcpMonitorQualityQueryResult(code int, data models.TcpQueryQualitySliceMonitor, meta *schemas.Pagination) *TcpMonitorQualityQueryResult {
	return &TcpMonitorQualityQueryResult{
		Code: code,
		Msg:  ResultText(code),
		Meta: *meta,
		Data: data.Data,
	}
}

// 中间表
func NewMonitorQueryInResult(code int, data models.QueryIntermediateSliceMonitor, meta *schemas.Pagination) *MonitorQueryInResult {
	return &MonitorQueryInResult{
		Code: code,
		Msg:  ResultText(code),
		Meta: *meta,
		Data: data.Data,
	}
}
