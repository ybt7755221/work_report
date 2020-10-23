package service

import (
	et "work_report/entities"
	"work_report/models"
)

type WrUsersService struct {
}
/**
 * 根据多条件查询数据-单条
 */
func (c *WrUsersService) FindOne(conditions *et.WrUsers) (*et.WrUsers, error) {
	wrUsersModel := models.WrUsersModel{}
	return wrUsersModel.FindOne(conditions)
}
/**
 * 根据多条件查询数据
 */
func (c *WrUsersService) Find(conditions *et.WrUsers, pagination *et.Pagination) ([]et.WrUsers, error) {
	wrUsersModel := models.WrUsersModel{}
	wrUsersPage, err := wrUsersModel.Find(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrUsersPage, nil
}

/**
 * 根据多条件查询数据-分页
 */
func (c *WrUsersService) FindPaging(conditions *et.WrUsers, pagination *et.Pagination) (*et.WrUsersPageDao, error) {
	wrUsersModel := models.WrUsersModel{}
	wrUsersPage, err := wrUsersModel.FindPaging(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrUsersPage, nil
}

func (c *WrUsersService) FindById(id int) (*et.WrUsers, error) {
	wrUsersModel := models.WrUsersModel{}
	return wrUsersModel.GetById(id)
}

func (c *WrUsersService) Insert(wrUsers *et.WrUsers) (err error) {
	wrUsersModel := models.WrUsersModel{}
	err = wrUsersModel.Insert(wrUsers)
	if err != nil {
		return err
	}
	return nil
}

func (c *WrUsersService) UpdateById(id int, wrUsers *et.WrUsers) (has int64, err error) {
	wrUsersModel := models.WrUsersModel{}
	has, err = wrUsersModel.UpdateById(id, wrUsers)
	return
}