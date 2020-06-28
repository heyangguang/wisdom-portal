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

// 监控数据表 kafka
type MonitorKafkaMigrate struct {
	models.BaseMigrate
	Ip     string    `gorm:"not null;comment:'IP地址'" json:"ip"`
	Port   string    `gorm:"not null;comment:'端口'" json:"port"`
	Name   string    `gorm:"not null;comment:'节点名'" json:"name"`
	Status bool      `gorm:"not null;comment:'状态'" json:"status"`
	Time   time.Time `gorm:"not null;comment:'采集时间'" json:"time"`
	Tag    string    `gorm:"not null;comment:'分类'" json:"tag"`
}

func (t *MonitorKafkaMigrate) TableName() string {
	return "monitor_kafka"
}

// 监控数据表 kubernetes
type MonitorKubernetesMigrate struct {
	models.BaseMigrate
	Ip     string    `gorm:"not null;comment:'IP地址'" json:"ip"`
	Port   string    `gorm:"not null;comment:'端口'" json:"port"`
	Name   string    `gorm:"not null;comment:'节点名'" json:"name"`
	Status bool      `gorm:"not null;comment:'状态'" json:"status"`
	Time   time.Time `gorm:"not null;comment:'采集时间'" json:"time"`
	Tag    string    `gorm:"not null;comment:'分类'" json:"tag"`
}

func (t *MonitorKubernetesMigrate) TableName() string {
	return "monitor_kubernetes"
}

// 监控中间表 intermediate
type MonitorIntermediateMigrate struct {
	models.BaseMigrate
	Status bool   `gorm:"not null;comment:'状态'" json:"status"`
	Name   string `gorm:"not null;comment:'名字'" json:"name"`
	// tag用来区分是上传upload 还是中间表程序core
	Tag      string    `gorm:"not null;comment:'分类'" json:"tag"`
	Count    int       `gorm:"not null;comment:'条数'" json:"count"`
	Time     time.Time `gorm:"not null;comment:'时间'" json:"time"`
	Describe string    `gorm:"comment:'详情'" json:"describe"`
}

func (t *MonitorIntermediateMigrate) TableName() string {
	return "monitor_intermediate"
}

// accessLog不同站点执行点监控
type MonitorAccessLogMigrate struct {
	models.BaseMigrate
	Site string    `gorm:"not null;comment:'区域站点'" json:"site"`
	Tag  string    `gorm:"not null;comment:'标签exe_time代表执行点时间，预留'" json:"tag"`
	Time time.Time `gorm:"not null;comment:'执行点时间'" json:"time"`
}

func (t *MonitorAccessLogMigrate) TableName() string {
	return "monitor_accesslog"
}
