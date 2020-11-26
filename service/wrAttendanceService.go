package service

import (
	et "work_report/entities"
	"work_report/models"
)

type WrAttendanceService struct {
}
/**
 * 根据多条件查询数据-单条
 */
func (c *WrAttendanceService) FindOne(conditions *et.WrAttendance) (*et.WrAttendance, error) {
	wrAttendanceModel := models.WrAttendanceModel{}
	return wrAttendanceModel.FindOne(conditions)
}
/**
 * 根据多条件查询数据
 */
func (c *WrAttendanceService) Find(conditions *et.WrAttendance, pagination *et.Pagination) ([]et.WrAttendance, error) {
	wrAttendanceModel := models.WrAttendanceModel{}
	wrAttendancePage, err := wrAttendanceModel.Find(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrAttendancePage, nil
}

/**
 * 根据多条件查询数据-分页
 */
func (c *WrAttendanceService) FindPaging(conditions *et.WrAttendance, pagination *et.Pagination) (*et.WrAttendancePageDao, error) {
	wrAttendanceModel := models.WrAttendanceModel{}
	wrAttendancePage, err := wrAttendanceModel.FindPaging(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrAttendancePage, nil
}

func (c *WrAttendanceService) FindById(id int) (*et.WrAttendance, error) {
	wrAttendanceModel := models.WrAttendanceModel{}
	return wrAttendanceModel.GetById(id)
}

func (c *WrAttendanceService) Insert(wrAttendance *et.WrAttendance) (err error) {
	wrAttendanceModel := models.WrAttendanceModel{}
	err = wrAttendanceModel.Insert(wrAttendance)
	if err != nil {
		return err
	}
	return nil
}

func (c *WrAttendanceService) UpdateById(id int, wrAttendance *et.WrAttendance) (has int64, err error) {
	wrAttendanceModel := models.WrAttendanceModel{}
	has, err = wrAttendanceModel.UpdateById(id, wrAttendance)
	return
}