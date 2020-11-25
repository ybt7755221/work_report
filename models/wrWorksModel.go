package models

import (
	"errors"
	"fmt"
	"strings"
	. "work_report/entities"
	DB "work_report/libraries/database"
)

type WrWorksModel struct {
}

//查找多条数据
func (u *WrWorksModel) Find(conditions *WrWorks, pagination *Pagination) ([]WrWorks, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrWorksPage := new(WrWorksPageDao)
	wrWorksPage.PageNum = pagination.PageNum
	wrWorksPage.PageSize = pagination.PageSize
	//排序
	sort := pageinfo["sort"].(map[string]string)
	if len(sort) > 0 {
		for key, val := range sort {
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			} else {
				dbC = dbC.Desc(key)
			}
		}
	}
	//执行查找
	err := dbC.Find(&wrWorksPage.List, conditions)
	return wrWorksPage.List, err
}

//查找多条数据
func (u *WrWorksModel) FindByWeekly(conditions *WrWorks, pagination *Pagination, startTime string, endTime string) ([]WrWorks, error) {
	dbConn := DB.GetDB(Gin).Where("created >= ?", startTime).Where("created <= ?", endTime)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrWorksPage := new(WrWorksPageDao)
	wrWorksPage.PageNum = pagination.PageNum
	wrWorksPage.PageSize = pagination.PageSize
	//排序
	sort := pageinfo["sort"].(map[string]string)
	if len(sort) > 0 {
		for key, val := range sort {
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			} else {
				dbC = dbC.Desc(key)
			}
		}
	}
	//执行查找
	err := dbC.Find(&wrWorksPage.List, conditions)
	return wrWorksPage.List, err
}

//查找多条数据-分页
func (u *WrWorksModel) FindPaging(conditions *WrWorks, pagination *Pagination) (*WrWorksPageDao, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrWorksPage := new(WrWorksPageDao)
	wrWorksPage.PageNum = pagination.PageNum
	wrWorksPage.PageSize = pagination.PageSize
	//排序
	sort := pageinfo["sort"].(map[string]string)
	if len(sort) > 0 {
		for key, val := range sort {
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			} else {
				dbC = dbC.Desc(key)
			}
		}
	}
	//执行查找
	err := dbC.Find(&wrWorksPage.List, conditions)
	total, err := dbC.Count(conditions)
	if err == nil {
		wrWorksPage.Total = total
	}
	return wrWorksPage, err
}

//根据id查找单条数据
func (u *WrWorksModel) GetById(id int) (*WrWorks, error) {
	fmt.Println(id)
	wrWorks := &WrWorks{Id: id}
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(wrWorks)
	defer dbConn.Close()
	return wrWorks, err
}

//插入
func (u *WrWorksModel) Insert(wrWorks *WrWorks) (err error) {
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Insert(wrWorks)
	defer dbConn.Close()
	if err != nil {
		return err
	}
	if affected < 1 {
		err = errors.New("插入影响行数: 0")
		return err
	}
	return err
}

//根据id更新
func (u *WrWorksModel) UpdateById(id int, wrWorks *WrWorks) (affected int64, err error) {
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Id(id).Update(wrWorks)
	defer dbConn.Close()
	return
}

func (u *WrWorksModel) DeleteById(id int, userId int) (int64, error) {
	dbConn := DB.GetDB(Gin)
	wrWorks := new(WrWorks)
	num, err := dbConn.Id(id).Where("user_id = ?", userId).Delete(wrWorks)
	return num, err
}
