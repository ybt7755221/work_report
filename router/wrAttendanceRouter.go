package router

import (
	. "work_report/controllers"

	"github.com/gin-gonic/gin"
)

func wrAttendanceRouter(router *gin.Engine) {
	wrAttendance := WrAttendanceController{}
	wrAttendanceR := router.Group("attendance")
	{
		wrAttendanceR.GET("/", wrAttendance.Find)
		wrAttendanceR.GET("/one", wrAttendance.FindOne)
		wrAttendanceR.GET("/page", wrAttendance.FindPaging)
		wrAttendanceR.POST("/", wrAttendance.Create)
		wrAttendanceR.GET("/find-by-id/:id", wrAttendance.FindById)
		wrAttendanceR.POST("/update-by-id", wrAttendance.UpdateById)
		wrAttendanceR.POST("/delete-by-id", wrAttendance.DeleteById)
	}
}
