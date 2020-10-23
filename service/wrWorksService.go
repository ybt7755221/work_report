package service

import (
	et "work_report/entities"
	"work_report/models"
)

type WrWorksService struct {
}
/**
 * 根据多条件查询数据
 */
func (c *WrWorksService) Find(conditions *et.WrWorks, pagination *et.Pagination) ([]et.WrWorks, error) {
	wrWorksModel := models.WrWorksModel{}
	wrWorksPage, err := wrWorksModel.Find(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrWorksPage, nil
}

/**
 * 根据多条件查询数据-分页
 */
func (c *WrWorksService) FindPaging(conditions *et.WrWorks, pagination *et.Pagination) (*et.WrWorksPageDao, error) {
	wrWorksModel := models.WrWorksModel{}
	wrWorksPage, err := wrWorksModel.FindPaging(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrWorksPage, nil
}

func (c *WrWorksService) FindByWeekly(conditions *et.WrWorks, pagination *et.Pagination, startTime string, endTime string) ([]et.WrWorks, error) {
	wrWorksModel := models.WrWorksModel{}
	wrWorksPage, err := wrWorksModel.FindByWeekly(conditions, pagination, startTime, endTime)
	if err != nil {
		return nil, err
	}
	return wrWorksPage, nil
}

func (c *WrWorksService) FindById(id int) (*et.WrWorks, error) {
	wrWorksModel := models.WrWorksModel{}
	return wrWorksModel.GetById(id)
}

func (c *WrWorksService) Insert(wrWorks *et.WrWorks) (err error) {
	wrWorksModel := models.WrWorksModel{}
	err = wrWorksModel.Insert(wrWorks)
	if err != nil {
		return err
	}
	return nil
}

func (c *WrWorksService) UpdateById(id int, wrWorks *et.WrWorks) (has int64, err error) {
	wrWorksModel := models.WrWorksModel{}
	has, err = wrWorksModel.UpdateById(id, wrWorks)
	return
}

func (c *WrWorksService) DeletedById(id int, userId int) (bool, error) {
	wrWorksModel := models.WrWorksModel{}
	has, err := wrWorksModel.DeleteById(id, userId)
	if err != nil || has == 0 {
		return false, err
	}
	return true, err
}