package router

import (
	. "work_report/controllers"
	"github.com/gin-gonic/gin"
)

func wrWorksRouter(router *gin.Engine) {
	wrWorks := WrWorksController{}
	wrWorksR := router.Group("work")
	{
		wrWorksR.GET("/", wrWorks.Find)
		wrWorksR.GET("/weekly", wrWorks.FindByWeekly)
		wrWorksR.POST("/delete", wrWorks.DeletedById)
		wrWorksR.GET("/page", wrWorks.FindPaging)
		wrWorksR.POST("/", wrWorks.Create)
		wrWorksR.GET("/find-by-id/:id", wrWorks.FindById)
		wrWorksR.PUT("/update-by-id", wrWorks.UpdateById)
	}
}