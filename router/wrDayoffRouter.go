package router

import (
	. "work_report/controllers"

	"github.com/gin-gonic/gin"
)

func wrDayoffRouter(router *gin.Engine) {
	wrDayoff := WrDayoffController{}
	wrDayoffR := router.Group("dayoff")
	{
		wrDayoffR.GET("/", wrDayoff.Find)
		wrDayoffR.GET("/one", wrDayoff.FindOne)
		wrDayoffR.GET("/page", wrDayoff.FindPaging)
		wrDayoffR.POST("/", wrDayoff.Create)
		wrDayoffR.GET("/find-by-id/:id", wrDayoff.FindById)
		wrDayoffR.POST("/update-by-id", wrDayoff.UpdateById)
		wrDayoffR.POST("/delete-by-attendance-id", wrDayoff.DeleteByAttendanceId)
		wrDayoffR.POST("/delete-by-id", wrDayoff.DeleteById)
	}
}
