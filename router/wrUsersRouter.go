package router

import (
	. "work_report/controllers"
	"github.com/gin-gonic/gin"
)

func wrUsersRouter(router *gin.Engine) {
	wrUsers := WrUsersController{}
	router.POST("check", wrUsers.Check)
	wrUsersR := router.Group("user")
	{
		wrUsersR.GET("/", wrUsers.Find)
		wrUsersR.GET("/one", wrUsers.FindOne)
		wrUsersR.GET("/page", wrUsers.FindPaging)
		wrUsersR.POST("/", wrUsers.Create)
		wrUsersR.GET("/find-by-id/:id", wrUsers.FindById)
		wrUsersR.PUT("/update-by-id", wrUsers.UpdateById)
	}
}