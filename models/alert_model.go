package models

import (
	"fmt"
	"strings"
	"time"
	"wisdom-portal/schemas"
	"wisdom-portal/wisdom-portal/logger"
)

type Alert struct {
	BaseModel
	Status      string    `gorm:"comment:'告警状态'" json:"status"`
	Instance    string    `gorm:"comment:'实例'" json:"instance"`
	Description string    `gorm:"comment:'告警描述'" json:"description"`
	Summary     string    `gorm:"comment:'告警摘要'" json:"summary"`
	StartAt     time.Time `gorm:"comment:'告警开始时间'" json:"start_at"`
	EndAt       time.Time `gorm:"comment:'告警结束时间'" json:"end_at"`
	AlertName   string    `gorm:"comment:'告警类型'" json:"alert_name"`
	Level       string    `gorm:"comment:'告警级别'" json:"level"`
}

type QuerySliceAlert struct {
	schemas.BasePagination
	Data []QueryAlert        `form:"-"`
	Meta *schemas.Pagination `form:"-"`
}

type QueryAlert struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
	Instance    string    `json:"instance"`
	Description string    `json:"description"`
	Summary     string    `json:"summary"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	AlertName   string    `json:"alert_name"`
	Level       string    `json:"level"`
}

// 创建数据不执行数据库操作
func (a *Alert) CreateAlert(am *schemas.AlertManagerWebHook) error {
	if err := a.formatAlert(am); err != nil {
		return err
	}
	return nil
}

// 更新数据
// 更新监控的状态
func (a *Alert) UpdateAlert() error {
	updateAlert := Alert{
		Status: "resolved",
		EndAt:  time.Now(),
	}
	if err := DB.Model(&a).Updates(&updateAlert).Error; err != nil {
		return err
	}
	return nil
}

// 查询数据
func (a *QuerySliceAlert) QueryAlert(startNum, endNum int) error {
	var queryAlert []QueryAlert
	if err := DB.Table("alert").Limit(endNum).Offset(startNum).Find(&queryAlert).Error; err != nil {
		return err
	}
	a.Data = queryAlert
	return nil
}

// 获取总数
func (a *QuerySliceAlert) CountNum() (num int, err error) {
	if err := DB.Table("alert").Model(QueryAlert{}).Count(&num).Error; err != nil {
		return 0, err
	}
	return num, nil
}

// 查询数据是否存在
func (a *Alert) QueryId(id string) error {
	if err := DB.Where("id = ?", id).Take(&a).Error; err != nil {
		return err
	}
	return nil
}

// 插入数据执行数据库操作
func (a *Alert) insertAlert() error {
	// Create方法会返回ID，导致ID重复，所以这里赋值0就不会重复
	a.ID = 0
	if err := DB.Create(a).Error; err != nil {
		logger.Error("insertAlert插入数据失败, err" + err.Error())
		return err
	}
	return nil
}

// 数据转换
func (a *Alert) formatAlert(am *schemas.AlertManagerWebHook) error {
	var err error
	for _, alertObj := range am.Alerts {
		a.Status = alertObj.Status
		a.Instance = alertObj.Labels.Instance
		a.Description = alertObj.Annotations.Description
		a.Summary = alertObj.Annotations.Summary
		a.StartAt = a.formatAt(alertObj.StartAt)
		a.EndAt, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
		a.AlertName = alertObj.Labels.AlertName
		a.Level = alertObj.Labels.Level
		err = a.insertAlert()
		if err != nil {
			logger.Error(fmt.Sprintf("formatAlert error: %v", err))
			break
		}
	}
	logger.Debug(fmt.Sprintf("%v", err))
	logger.Debug("formatAlert结束")
	return err
}

// 时间字符串转换
// string 2020-05-27T11:23:37.380201056Z
// time.Time 2020-05-27 11:23:37.380201056
func (a *Alert) formatAt(t string) time.Time {
	var at time.Time
	timeTemplate := "2006-01-02 15:04:05"
	t = strings.Replace(t, "T", " ", 1)
	t = strings.Replace(t, "Z", "", 1)
	logger.Debug(t)
	at, _ = time.ParseInLocation(timeTemplate, t, time.Local)
	// 时区转换 +8 转东八区
	h, _ := time.ParseDuration("1h")
	h8 := at.Add(h * 8)
	return h8
}
