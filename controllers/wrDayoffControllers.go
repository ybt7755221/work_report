package controllers

import (
	et "work_report/entities"
	"work_report/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
type WrDayoffController struct {
	serv *service.WrDayoffService
}
// @Tags user表操作
// @Summary 【GetOne】根据条件获取单条数据
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Success 200 {object} SgrResp
// @Router /user/one [get]
func (c *WrDayoffController) FindOne(ctx *gin.Context) {
	wrDayoff := new(et.WrDayoff)
	getParamsNew(ctx, wrDayoff)
	res, err := c.serv.FindOne(wrDayoff)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		resSuccess(ctx, res)
	}
}
// @Tags dayoff表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序 {\"id\":\"desc\"}"
// @Success 200 {object} SgrResp
// @Router /dayoff [get]
func (c *WrDayoffController) Find(ctx *gin.Context) {
	wrDayoff := new(et.WrDayoff)
	getParamsNew(ctx, wrDayoff)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrDayoffList, err := c.serv.Find(wrDayoff, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrDayoffList)
}
// @Tags dayoff表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序 {\"id\":\"desc\"}"
// @Success 200 {object} SgrResp
// @Router /dayoff/page [get]
func (c *WrDayoffController) FindPaging(ctx *gin.Context) {
	wrDayoff := new(et.WrDayoff)
	getParamsNew(ctx, wrDayoff)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrDayoffList, err := c.serv.FindPaging(wrDayoff, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrDayoffList)
}
// @Tags dayoff表操作
// @Summary 【GetOne】根据id获取信息
// @Description 根据id获取信息
// @Accept html
// @Produce json
// @Param   id		path	string 	false	"主键id"
// @Success 200 {object} SgrResp
// @Router /dayoff/find-by-id/{id} [get]
func (c *WrDayoffController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	wrDayoff, err := c.serv.FindById(id)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		resSuccess(ctx, wrDayoff)
	}
}
// @Tags dayoff表操作
// @Summary 【create】创建dayoff信息
// @Description 创建dayoff信息
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} SgrResp
// @Router /dayoff [post]
func (c *WrDayoffController) Create(ctx *gin.Context) {
	wrDayoff := new(et.WrDayoff)
	getPostStructData(ctx, wrDayoff)
	if err := c.serv.Insert(wrDayoff); err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrDayoff)
}
// @Tags dayoff表操作
// @Summary 【update】根据id更新数据
// @Description 根据id更新数据
// @Accept x-www-form-urlencoded
// @Produce json
// @Param   id	body	string 	true	"主键更新依据此id"
// @Success 200 {object} SgrResp
// @Router /dayoff/update-by-id [put]
func (c * WrDayoffController) UpdateById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	wrDayoff := new(et.WrDayoff)
	getPostStructData(ctx, wrDayoff)
	has, err := c.serv.UpdateById(id, wrDayoff)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		if has == 0 {
			resError(ctx, et.EntityFailure, "影响行数0")
		}else{
			resSuccess(ctx, gin.H{
				"update_count":has,
			})
		}
	}
}