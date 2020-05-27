package schemas

import (
	"fmt"
	"go.uber.org/zap"
	"wisdom-portal/wisdom-portal/logger"
)

//const (
//	PageSize = 4
//)

// 查询
type BasePagination struct {
	Page     int `form:"page" validate:"required" label:"page"`
	PageSize int `form:"page_size" validate:"required" label:"page_size"`
}

// 分页器
type Pagination struct {
	HasPre   bool   `json:"has_pre"`
	HasNext  bool   `json:"has_next"`
	NumCount int    `json:"num_count"`
	ShowPage string `json:"show_page"`
}

// 生成分页数据结构
func NewPagination(num int) *Pagination {
	pageObj := new(Pagination)

	// 获取总条数
	pageObj.NumCount = num

	logger.Debug("Pagination", zap.Any("pageObj", *pageObj))
	return pageObj
}

// 处理分页数据
func (pagination *Pagination) PaginationStint(page, pageSize int) (startNum, endNum int) {
	// 总页数
	allPageNum := (pagination.NumCount-1)/pageSize + 1

	if page <= 1 {
		page = 1
		pagination.HasNext = true
	}

	if page >= allPageNum {
		page = allPageNum
		pagination.HasPre = true
	}

	if pagination.NumCount == 0 {
		pagination.HasNext = false
		pagination.HasPre = false
	}

	// (page - 1)*PageSize  page * PageSize
	// 1
	// 0*4 1 = 0 4
	// 2
	// 1*4 2*4 = 4 8
	startNum = (page - 1) * pageSize
	endNum = page * pageSize

	// 取模
	//result := pagination.NumCount % pageSize
	//fmt.Println(result)
	if allPageNum == page {
		// 计算数字差
		//diffNum := pageSize - result
		//endNum = (page * pageSize) - diffNum
		endNum = pagination.NumCount
	}
	pagination.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)
	return
}
