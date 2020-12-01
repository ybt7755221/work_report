package models

import (
	"errors"
	"fmt"
	"strings"
	. "work_report/entities"
	DB "work_report/libraries/database"
)

type WrDayoffModel struct {
}

//查找单条数据
func (u *WrDayoffModel) FindOne(conditions *WrDayoff) (*WrDayoff, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	has, err := dbConn.Get(conditions)
	if has {
		return conditions, err
	} else {
		return nil, err
	}
}

//查找多条数据
func (u *WrDayoffModel) Find(conditions *WrDayoff, pagination *Pagination) ([]WrDayoff, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrDayoffPage := new(WrDayoffPageDao)
	wrDayoffPage.PageNum = pagination.PageNum
	wrDayoffPage.PageSize = pagination.PageSize
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
	err := dbC.Find(&wrDayoffPage.List, conditions)
	return wrDayoffPage.List, err
}

//查找多条数据-分页
func (u *WrDayoffModel) FindPaging(conditions *WrDayoff, pagination *Pagination) (*WrDayoffPageDao, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	wrDayoffPage := new(WrDayoffPageDao)
	wrDayoffPage.PageNum = pagination.PageNum
	wrDayoffPage.PageSize = pagination.PageSize
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
	err := dbC.Find(&wrDayoffPage.List, conditions)
	total, err := dbC.Count(conditions)
	if err == nil {
		wrDayoffPage.Total = total
	}
	return wrDayoffPage, err
}

//根据id查找单条数据
func (u *WrDayoffModel) GetById(id int) (*WrDayoff, error) {
	fmt.Println(id)
	wrDayoff := &WrDayoff{Id: id}
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(wrDayoff)
	defer dbConn.Close()
	return wrDayoff, err
}

//插入
func (u *WrDayoffModel) Insert(wrDayoff *WrDayoff) (err error) {
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Insert(wrDayoff)
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
func (u *WrDayoffModel) UpdateById(id int, wrDayoff *WrDayoff) (affected int64, err error) {
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Id(id).Update(wrDayoff)
	defer dbConn.Close()
	return
}

func (u *WrDayoffModel) DeleteById(id int, userId int) (int64, error) {
	dbConn := DB.GetDB(Gin)
	wrDayoff := new(WrDayoff)
	num, err := dbConn.Id(id).Where("user_id = ?", userId).Delete(wrDayoff)
	return num, err
}

func (u *WrDayoffModel) DeleteByAttendanceId(attendanceId int) (int64, error) {
	dbConn := DB.GetDB(Gin)
	wrDayoff := new(WrDayoff)
	num, err := dbConn.Where("attendance_id = ?", attendanceId).Delete(wrDayoff)
	return num, err
}
