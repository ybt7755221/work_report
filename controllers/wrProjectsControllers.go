package controllers

import (
	et "work_report/entities"
	"work_report/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
type WrProjectsController struct {
	serv *service.WrProjectsService
}
// @Tags project表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序 {\"id\":\"desc\"}"
// @Success 200 {object} SgrResp
// @Router /project [get]
func (c *WrProjectsController) Find(ctx *gin.Context) {
	wrProjects := new(et.WrProjects)
	getParamsNew(ctx, wrProjects)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrProjectsList, err := c.serv.Find(wrProjects, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrProjectsList)
}
// @Tags project表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序 {\"id\":\"desc\"}"
// @Success 200 {object} SgrResp
// @Router /project/page [get]
func (c *WrProjectsController) FindPaging(ctx *gin.Context) {
	wrProjects := new(et.WrProjects)
	getParamsNew(ctx, wrProjects)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrProjectsList, err := c.serv.FindPaging(wrProjects, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrProjectsList)
}
// @Tags project表操作
// @Summary 【GetOne】根据id获取信息
// @Description 根据id获取信息
// @Accept html
// @Produce json
// @Param   id		path	string 	false	"主键id"
// @Success 200 {object} SgrResp
// @Router /project/find-by-id/{id} [get]
func (c *WrProjectsController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	wrProjects, err := c.serv.FindById(id)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		resSuccess(ctx, wrProjects)
	}
}
// @Tags project表操作
// @Summary 【create】创建project信息
// @Description 创建project信息
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} SgrResp
// @Router /project [post]
func (c *WrProjectsController) Create(ctx *gin.Context) {
	wrProjects := new(et.WrProjects)
	getPostStructData(ctx, wrProjects)
	if err := c.serv.Insert(wrProjects); err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrProjects)
}
// @Tags project表操作
// @Summary 【update】根据id更新数据
// @Description 根据id更新数据
// @Accept x-www-form-urlencoded
// @Produce json
// @Param   id	body	string 	true	"主键更新依据此id"
// @Success 200 {object} SgrResp
// @Router /project/update-by-id [put]
func (c * WrProjectsController) UpdateById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	wrProjects := new(et.WrProjects)
	getPostStructData(ctx, wrProjects)
	has, err := c.serv.UpdateById(id, wrProjects)
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