package migrate

import (
	"time"
	"wisdom-portal/models"
)

// 告警数据表
type AlertMigrate struct {
	models.BaseMigrate
	Status      string    `gorm:"comment:'告警状态'" json:"status"`
	Instance    string    `gorm:"comment:'实例'" json:"instance"`
	Description string    `gorm:"comment:'告警描述'" json:"description"`
	Summary     string    `gorm:"comment:'告警摘要'" json:"summary"`
	StartAt     time.Time `gorm:"comment:'告警开始时间'" json:"start_at"`
	EndAt       time.Time `gorm:"comment:'告警结束时间'" json:"end_at"`
	AlertName   string    `gorm:"comment:'告警类型'" json:"alert_name"`
	Level       string    `gorm:"comment:'告警级别'" json:"level"`
}

func (t *AlertMigrate) TableName() string {
	return "alert"
}
