package router

import (
	. "work_report/controllers"
	"github.com/gin-gonic/gin"
)

func wrProjectsRouter(router *gin.Engine) {
	wrProjects := WrProjectsController{}
	wrProjectsR := router.Group("project")
	{
		wrProjectsR.GET("/", wrProjects.Find)
		wrProjectsR.GET("/page", wrProjects.FindPaging)
		wrProjectsR.POST("/", wrProjects.Create)
		wrProjectsR.GET("/find-by-id/:id", wrProjects.FindById)
		wrProjectsR.PUT("/update-by-id", wrProjects.UpdateById)
	}
}