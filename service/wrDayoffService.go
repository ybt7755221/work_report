package service

import (
	et "work_report/entities"
	"work_report/models"
)

type WrDayoffService struct {
}
/**
 * 根据多条件查询数据-单条
 */
func (c *WrDayoffService) FindOne(conditions *et.WrDayoff) (*et.WrDayoff, error) {
	wrDayoffModel := models.WrDayoffModel{}
	return wrDayoffModel.FindOne(conditions)
}
/**
 * 根据多条件查询数据
 */
func (c *WrDayoffService) Find(conditions *et.WrDayoff, pagination *et.Pagination) ([]et.WrDayoff, error) {
	wrDayoffModel := models.WrDayoffModel{}
	wrDayoffPage, err := wrDayoffModel.Find(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrDayoffPage, nil
}

/**
 * 根据多条件查询数据-分页
 */
func (c *WrDayoffService) FindPaging(conditions *et.WrDayoff, pagination *et.Pagination) (*et.WrDayoffPageDao, error) {
	wrDayoffModel := models.WrDayoffModel{}
	wrDayoffPage, err := wrDayoffModel.FindPaging(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrDayoffPage, nil
}

func (c *WrDayoffService) FindById(id int) (*et.WrDayoff, error) {
	wrDayoffModel := models.WrDayoffModel{}
	return wrDayoffModel.GetById(id)
}

func (c *WrDayoffService) Insert(wrDayoff *et.WrDayoff) (err error) {
	wrDayoffModel := models.WrDayoffModel{}
	err = wrDayoffModel.Insert(wrDayoff)
	if err != nil {
		return err
	}
	return nil
}

func (c *WrDayoffService) UpdateById(id int, wrDayoff *et.WrDayoff) (has int64, err error) {
	wrDayoffModel := models.WrDayoffModel{}
	has, err = wrDayoffModel.UpdateById(id, wrDayoff)
	return
}