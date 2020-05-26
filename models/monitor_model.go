package models

import (
	"time"
	"wisdom-portal/schemas"
	"wisdom-portal/wisdom-portal/logger"
)

type Monitor struct {
	BaseModel
	Ip     string    `gorm:"not null;comment:'IP地址'" json:"ip"`
	Port   string    `gorm:"not null;comment:'端口'" json:"port"`
	Name   string    `gorm:"not null;comment:'节点名'" json:"name"`
	Status bool      `gorm:"not null;comment:'状态'" json:"status"`
	Time   time.Time `gorm:"not null;comment:'采集时间'" json:"time"`
	Tag    string    `gorm:"not null;comment:'分类'" json:"tag"`
}

type QuerySliceMonitor struct {
	AppTag   string                      `form:"app_tag" validate:"required,ValidationAppTagFormat" label:"app_tag"`
	Num      int                         `form:"num" validate:"required,ValidationNumFormat" label:"num"`
	Page     int                         `form:"page" validate:"required" label:"page"`
	PageSize int                         `form:"page_size" validate:"required" label:"page_size"`
	Data     []map[string][]QueryMonitor `form:"-"`
	Meta     *schemas.Pagination         `form:"-"`
}

type QueryMonitor struct {
	Ip     string    `json:"ip"`
	Port   string    `json:"port"`
	Name   string    `json:"name"`
	Status bool      `json:"status"`
	Time   time.Time `json:"time"`
	Tag    string    `json:"tag"`
}

type TagGroupBy struct {
	Name string
}

// 插入监控数据
func (m *Monitor) CreateMonitor() error {
	switch m.Tag {
	case "MySQL":
		if err := DB.Table("monitor_mysql").Create(&m).Error; err != nil {
			logger.Error("CreateMonitor monitor_mysql 插入数据失败, err:" + err.Error())
			return err
		}
	case "ElasticSearch":
		if err := DB.Table("monitor_elasticsearch").Create(&m).Error; err != nil {
			logger.Error("CreateMonitor monitor_elasticsearch 插入数据失败, err:" + err.Error())
			return err
		}
	}
	return nil
}

// 查询服务监控状态
func (querySlice *QuerySliceMonitor) QueryMonitor(startNum, endNum int) error {
	if err := querySlice.selectQueryApp(startNum, endNum); err != nil {
		return err
	}
	return nil
}

// 选择不同的APP查询
func (querySlice *QuerySliceMonitor) selectQueryApp(startNum, endNum int) error {
	//var tagGroupBy []TagGroupBy
	//if err := DB.Table(querySlice.castTableName()).Select("name").Group("name").Scan(&tagGroupBy).Error; err != nil {
	//	logger.Error("selectQueryApp 查询数据失败, err:" + err.Error())
	//	return err
	//}
	tagGroupBy, _, err := querySlice.CountNum()
	if err != nil {
		logger.Error("selectQueryApp 查询数据失败, err:" + err.Error())
		return err
	}
	for _, tagName := range tagGroupBy[startNum:endNum] {
		var queryMonitorObj []QueryMonitor
		if err := DB.Table(querySlice.castTableName()).Where("name = ?", tagName.Name).Order("time desc").Limit(querySlice.Num).Find(&queryMonitorObj).Error; err != nil {
			logger.Error("selectQueryApp 查询数据失败, err:" + err.Error())
			return err
		}
		querySlice.Data = append(querySlice.Data, map[string][]QueryMonitor{tagName.Name: queryMonitorObj})
	}
	return nil
}

// APP对应数据表关系
func (querySlice *QuerySliceMonitor) castTableName() string {
	switch querySlice.AppTag {
	case "MySQL":
		return "monitor_mysql"
	case "ElasticSearch":
		return "monitor_elasticsearch"
	}
	return ""
}

// 数据总数
func (querySlice *QuerySliceMonitor) CountNum() (tagGroupBy []TagGroupBy, num int, err error) {
	if err := DB.Table(querySlice.castTableName()).Select("name").Group("name").Scan(&tagGroupBy).Error; err != nil {
		logger.Error("selectQueryApp 查询数据失败, err:" + err.Error())
		return nil, 0, err
	}
	return tagGroupBy, len(tagGroupBy), nil
}
