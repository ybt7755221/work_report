package models

import (
	. "work_report/entities"
	DB "work_report/libraries/database"
	"errors"
	"fmt"
	"strings"
)

type WrAttendanceModel struct {
}

//查找单条数据
func(u *WrAttendanceModel) FindOne(conditions *WrAttendance) (*WrAttendance, error) {
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
func(u *WrAttendanceModel) Find(conditions *WrAttendance, pagination *Pagination )  ([]WrAttendance, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrAttendancePage := new(WrAttendancePageDao)
	wrAttendancePage.PageNum = pagination.PageNum
	wrAttendancePage.PageSize = pagination.PageSize
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
	err := dbC.Find(&wrAttendancePage.List, conditions)
	return wrAttendancePage.List, err
}

//查找多条数据-分页
func(u *WrAttendanceModel) FindPaging(conditions *WrAttendance, pagination *Pagination )  (*WrAttendancePageDao, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrAttendancePage := new(WrAttendancePageDao)
	wrAttendancePage.PageNum = pagination.PageNum
	wrAttendancePage.PageSize = pagination.PageSize
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
	err := dbC.Find(&wrAttendancePage.List, conditions)
	total, err := dbC.Count(conditions)
	if err == nil {
		wrAttendancePage.Total = total
	}
	return wrAttendancePage, err
}
//根据id查找单条数据
func (u *WrAttendanceModel) GetById(id int) (*WrAttendance, error) {
	fmt.Println(id)
	wrAttendance := &WrAttendance{Id: id}
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(wrAttendance)
	defer dbConn.Close()
	return wrAttendance, err
}
//插入
func (u *WrAttendanceModel) Insert(wrAttendance *WrAttendance) (err error) {
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Insert(wrAttendance)
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
func (u *WrAttendanceModel) UpdateById(id int, wrAttendance *WrAttendance) (affected int64, err error) {
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Id(id).Update(wrAttendance)
	defer dbConn.Close()
	return
}