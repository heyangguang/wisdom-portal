package models

import (
	"fmt"
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

// 状态监控
type TcpQuerySliceMonitor struct {
	schemas.BasePagination
	AppTag string                         `form:"app_tag" validate:"required,ValidationAppTagFormat" label:"app_tag"`
	Num    int                            `form:"num" validate:"required,ValidationNumFormat" label:"num"`
	Data   []map[string][]TcpQueryMonitor `form:"-"`
	Meta   *schemas.Pagination            `form:"-"`
}

// 状态监控实例
type TcpQueryMonitor struct {
	Ip     string    `json:"ip"`
	Port   string    `json:"port"`
	Name   string    `json:"name"`
	Status bool      `json:"status"`
	Time   time.Time `json:"time"`
	Tag    string    `json:"tag"`
}

// 质量监控
type TcpQueryQualitySliceMonitor struct {
	schemas.BasePagination
	AppTag   string                   `form:"app_tag" validate:"required,ValidationAppTagFormat" label:"app_tag"`
	Interval string                   `form:"interval" validate:"required,ValidationIntervalFormat" label:"interval"`
	Data     []TcpQueryQualityMonitor `form:"-"`
	Meta     *schemas.Pagination      `form:"-"`
}

// 质量监控实例
type TcpQueryQualityMonitor struct {
	Name       string  `json:"name"`
	Ip         string  `json:"ip"`
	Port       string  `json:"port"`
	SuccessAvg float64 `json:"success_avg"`
}

// 分组
type TagGroupBy struct {
	Name string
	Ip   string
	Port string
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
func (q *TcpQuerySliceMonitor) QueryMonitor(startNum, endNum int) error {
	if err := q.selectQueryApp(startNum, endNum); err != nil {
		return err
	}
	return nil
}

// 查询服务质量状态
func (q *TcpQueryQualitySliceMonitor) QueryMonitor(startNum, endNum int) error {
	if err := q.selectQueryApp(startNum, endNum); err != nil {
		return err
	}
	return nil
}

// 选择不同的APP查询
func (q *TcpQuerySliceMonitor) selectQueryApp(startNum, endNum int) error {
	tagGroupBy, _, err := q.CountNum()
	if err != nil {
		logger.Error("selectQueryApp 查询数据失败, err:" + err.Error())
		return err
	}
	for _, tagName := range tagGroupBy[startNum:endNum] {
		var queryMonitorObj []TcpQueryMonitor
		if err := DB.Table(CastTableName(q.AppTag)).Where("name = ?", tagName.Name).Order("time desc").Limit(q.Num).Find(&queryMonitorObj).Error; err != nil {
			logger.Error("selectQueryApp 查询数据失败, err:" + err.Error())
			return err
		}
		logger.Debug(fmt.Sprintf("服务监控状态查询数据：%v", queryMonitorObj))
		q.Data = append(q.Data, map[string][]TcpQueryMonitor{tagName.Name: queryMonitorObj})
	}
	return nil
}

// 质量监控选择APP查询
func (q *TcpQueryQualitySliceMonitor) selectQueryApp(startNum, endNum int) error {
	tagGroupBy, _, err := q.CountNum()
	if err != nil {
		logger.Error("selectQueryApp 查询数据失败, err:" + err.Error())
		return err
	}
	timeNow := time.Now()
	for _, tagName := range tagGroupBy[startNum:endNum] {
		var queryQualityMonitorObj TcpQueryQualityMonitor
		m, _ := time.ParseDuration(fmt.Sprintf("-%sm", q.Interval))
		sql := fmt.Sprintf("SELECT name, ip, port, avg(status) as success_avg FROM `%s` "+
			"WHERE (created_at >= ?) and name = ? and ip = ? and port = ? GROUP BY `name`, `ip`, `port`", CastTableName(q.AppTag))
		if err := DB.Raw(sql, timeNow.Add(m), tagName.Name, tagName.Ip, tagName.Port).Scan(&queryQualityMonitorObj).Error; err != nil {
			logger.Error("selectQueryApp 查询数据失败, err:" + err.Error())
			return err
		}
		logger.Debug(fmt.Sprintf("服务监控质量查询数据：%v", queryQualityMonitorObj))
		q.Data = append(q.Data, queryQualityMonitorObj)
	}
	return nil
}

// 数据总数
func (q *TcpQuerySliceMonitor) CountNum() (tagGroupBy []TagGroupBy, num int, err error) {
	if err := DB.Table(CastTableName(q.AppTag)).Select("name").Group("name").Scan(&tagGroupBy).Error; err != nil {
		logger.Error("selectQueryApp 查询数据失败, err:" + err.Error())
		return nil, 0, err
	}
	return tagGroupBy, len(tagGroupBy), nil
}

// 质量监控数据总数
func (q *TcpQueryQualitySliceMonitor) CountNum() (tagGroupBy []TagGroupBy, num int, err error) {
	sql := fmt.Sprintf("SELECT name, ip, port FROM `%s` "+
		"WHERE (created_at >= ?) GROUP BY `name`, `ip`, `port`", CastTableName(q.AppTag))
	timeNow := time.Now()
	m, _ := time.ParseDuration(fmt.Sprintf("-%sm", q.Interval))
	if err := DB.Raw(sql, timeNow.Add(m)).Scan(&tagGroupBy).Error; err != nil {
		logger.Error("queryQualitySlice 查询数据失败, err:" + err.Error())
		return nil, 0, err
	}
	return tagGroupBy, len(tagGroupBy), nil
}

// APP对应数据表关系
func CastTableName(appTag string) string {
	switch appTag {
	case "MySQL":
		return "monitor_mysql"
	case "ElasticSearch":
		return "monitor_elasticsearch"
	}
	return ""
}
