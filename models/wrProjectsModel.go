package models

import (
	. "work_report/entities"
	DB "work_report/libraries/database"
	"errors"
	"fmt"
	"strings"
)

type WrProjectsModel struct {
}
//查找多条数据
func(u *WrProjectsModel) Find(conditions *WrProjects, pagination *Pagination )  ([]WrProjects, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrProjectsPage := new(WrProjectsPageDao)
	wrProjectsPage.PageNum = pagination.PageNum
	wrProjectsPage.PageSize = pagination.PageSize
	//排序
	sort := pageinfo["sort"].(map[string]string)
	if len(sort) > 0 {
		for key, val := range sort{
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			}else{
				dbC = dbC.Desc(key)
			}
		}
	}
	//执行查找
	err := dbC.Find(&wrProjectsPage.List, conditions)
	return wrProjectsPage.List, err
}

//查找多条数据-分页
func(u *WrProjectsModel) FindPaging(conditions *WrProjects, pagination *Pagination )  (*WrProjectsPageDao, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrProjectsPage := new(WrProjectsPageDao)
	wrProjectsPage.PageNum = pagination.PageNum
	wrProjectsPage.PageSize = pagination.PageSize
	//排序
	sort := pageinfo["sort"].(map[string]string)
	if len(sort) > 0 {
		for key, val := range sort{
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			}else{
				dbC = dbC.Desc(key)
			}
		}
	}
	//执行查找
	err := dbC.Find(&wrProjectsPage.List, conditions)
	total, err := dbC.Count(conditions)
	if err == nil {
		wrProjectsPage.Total = total
	}
	return wrProjectsPage, err
}
//根据id查找单条数据
func (u *WrProjectsModel) GetById(id int) (*WrProjects, error) {
	fmt.Println(id)
	wrProjects := &WrProjects{Id: id}
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(wrProjects)
	defer dbConn.Close()
	return wrProjects, err
}
//插入
func (u *WrProjectsModel) Insert(wrProjects *WrProjects) (err error) {
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Insert(wrProjects)
	defer dbConn.Close()
	if err != nil {
		return err
	}
	if affected < 1 {
		err = errors.New("插入影响行数: 0" )
		return err
	}
	return err
}
//根据id更新
func (u *WrProjectsModel) UpdateById(id int, wrProjects *WrProjects) (affected int64, err error) {
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Id(id).Update(wrProjects)
	defer dbConn.Close()
	return
}