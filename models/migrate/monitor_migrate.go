package migrate

import (
	"time"
	"wisdom-portal/models"
)

// 监控数据表 elasticSearch
type MonitorElasticSearchMigrate struct {
	models.BaseMigrate
	Ip     string    `gorm:"not null;comment:'IP地址'" json:"ip"`
	Port   string    `gorm:"not null;comment:'端口'" json:"port"`
	Name   string    `gorm:"not null;comment:'节点名'" json:"name"`
	Status bool      `gorm:"not null;comment:'状态'" json:"status"`
	Time   time.Time `gorm:"not null;comment:'采集时间'" json:"time"`
	Tag    string    `gorm:"not null;comment:'分类'" json:"tag"`
}

func (t *MonitorElasticSearchMigrate) TableName() string {
	return "monitor_elasticsearch"
}

// 监控数据表 mysql
type MonitorMySQLMigrate struct {
	models.BaseMigrate
	Ip     string    `gorm:"not null;comment:'IP地址'" json:"ip"`
	Port   string    `gorm:"not null;comment:'端口'" json:"port"`
	Name   string    `gorm:"not null;comment:'节点名'" json:"name"`
	Status bool      `gorm:"not null;comment:'状态'" json:"status"`
	Time   time.Time `gorm:"not null;comment:'采集时间'" json:"time"`
	Tag    string    `gorm:"not null;comment:'分类'" json:"tag"`
}

func (t *MonitorMySQLMigrate) TableName() string {
	return "monitor_mysql"
}
