package models

import (
	"encoding/json"
	"work_report/config"
	. "work_report/entities"

	"github.com/gin-gonic/gin"
)

const (
	Gin = config.Gin
)

func getPagingParams(pagination *Pagination) gin.H {
	var offset = 0
	if pagination.PageSize < 1 {
		pagination.PageSize = config.DefPageSize
	}
	if pagination.PageNum > 0 {
		offset = (pagination.PageNum - 1) * pagination.PageSize
	}
	sort := make(map[string]string, 0)
	if len(pagination.SortStr) > 0 {
		sort = getSort(pagination.SortStr)
	}
	pagingParams := gin.H{
		"offset": offset,
		"limit":  pagination.PageSize,
		"sort":   sort,
	}
	return pagingParams
}

func getSort(sortStr string) (sort map[string]string) {
	if err := json.Unmarshal([]byte(sortStr), &sort); err != nil {
		return nil
	}
	return
}
