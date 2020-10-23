package models

import (
	. "work_report/entities"
	DB "work_report/libraries/database"
	"errors"
	"fmt"
	"strings"
)

type WrUsersModel struct {
}

//查找单条数据
func(u *WrUsersModel) FindOne(conditions *WrUsers) (*WrUsers, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	has, err := dbConn.Get(conditions)
	if has {
		return conditions, err
	}else{
		return nil, err
	}
}
//查找多条数据
func(u *WrUsersModel) Find(conditions *WrUsers, pagination *Pagination )  ([]WrUsers, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrUsersPage := new(WrUsersPageDao)
	wrUsersPage.PageNum = pagination.PageNum
	wrUsersPage.PageSize = pagination.PageSize
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
	err := dbC.Find(&wrUsersPage.List, conditions)
	return wrUsersPage.List, err
}

//查找多条数据-分页
func(u *WrUsersModel) FindPaging(conditions *WrUsers, pagination *Pagination )  (*WrUsersPageDao, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrUsersPage := new(WrUsersPageDao)
	wrUsersPage.PageNum = pagination.PageNum
	wrUsersPage.PageSize = pagination.PageSize
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
	err := dbC.Find(&wrUsersPage.List, conditions)
	total, err := dbC.Count(conditions)
	if err == nil {
		wrUsersPage.Total = total
	}
	return wrUsersPage, err
}
//根据id查找单条数据
func (u *WrUsersModel) GetById(id int) (*WrUsers, error) {
	fmt.Println(id)
	wrUsers := &WrUsers{Id: id}
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(wrUsers)
	defer dbConn.Close()
	return wrUsers, err
}
//插入
func (u *WrUsersModel) Insert(wrUsers *WrUsers) (err error) {
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Insert(wrUsers)
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
func (u *WrUsersModel) UpdateById(id int, wrUsers *WrUsers) (affected int64, err error) {
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Id(id).Update(wrUsers)
	defer dbConn.Close()
	return
}