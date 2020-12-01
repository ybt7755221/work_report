package controllers

import (
	"strconv"
	et "work_report/entities"
	"work_report/service"

	"github.com/gin-gonic/gin"
)

type WrAttendanceController struct {
	serv *service.WrAttendanceService
}

// @Tags user表操作
// @Summary 【GetOne】根据条件获取单条数据
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Success 200 {object} SgrResp
// @Router /user/one [get]
func (c *WrAttendanceController) FindOne(ctx *gin.Context) {
	wrAttendance := new(et.WrAttendance)
	getParamsNew(ctx, wrAttendance)
	res, err := c.serv.FindOne(wrAttendance)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	} else {
		resSuccess(ctx, res)
	}
}

// @Tags attendance表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序 {\"id\":\"desc\"}"
// @Success 200 {object} SgrResp
// @Router /attendance [get]
func (c *WrAttendanceController) Find(ctx *gin.Context) {
	wrAttendance := new(et.WrAttendance)
	getParamsNew(ctx, wrAttendance)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrAttendanceList, err := c.serv.Find(wrAttendance, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrAttendanceList)
}

// @Tags attendance表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序 {\"id\":\"desc\"}"
// @Success 200 {object} SgrResp
// @Router /attendance/page [get]
func (c *WrAttendanceController) FindPaging(ctx *gin.Context) {
	wrAttendance := new(et.WrAttendance)
	getParamsNew(ctx, wrAttendance)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrAttendanceList, err := c.serv.FindPaging(wrAttendance, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrAttendanceList)
}

// @Tags attendance表操作
// @Summary 【GetOne】根据id获取信息
// @Description 根据id获取信息
// @Accept html
// @Produce json
// @Param   id		path	string 	false	"主键id"
// @Success 200 {object} SgrResp
// @Router /attendance/find-by-id/{id} [get]
func (c *WrAttendanceController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	wrAttendance, err := c.serv.FindById(id)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	} else {
		resSuccess(ctx, wrAttendance)
	}
}

// @Tags attendance表操作
// @Summary 【create】创建attendance信息
// @Description 创建attendance信息
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} SgrResp
// @Router /attendance [post]
func (c *WrAttendanceController) Create(ctx *gin.Context) {
	wrAttendance := new(et.WrAttendance)
	getPostStructData(ctx, wrAttendance)
	if err := c.serv.Insert(wrAttendance); err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrAttendance)
}

// @Tags attendance表操作
// @Summary 【update】根据id更新数据
// @Description 根据id更新数据
// @Accept x-www-form-urlencoded
// @Produce json
// @Param   id	body	string 	true	"主键更新依据此id"
// @Success 200 {object} SgrResp
// @Router /attendance/update-by-id [put]
func (c *WrAttendanceController) UpdateById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	wrAttendance := new(et.WrAttendance)
	getPostStructData(ctx, wrAttendance)
	has, err := c.serv.UpdateById(id, wrAttendance)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	} else {
		if has == 0 {
			resError(ctx, et.EntityFailure, "影响行数0")
		} else {
			resSuccess(ctx, gin.H{
				"update_count": has,
			})
		}
	}
}

func (c *WrAttendanceController) DeleteById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	userId, _ := strconv.Atoi(ctx.PostForm("user_id"))
	res, err := c.serv.DeletedById(id, userId)
	if res {
		resSuccess(ctx, gin.H{})
	} else {
		resError(ctx, et.EntityPanic, err.Error())
	}
	return
}
