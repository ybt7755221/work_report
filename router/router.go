package router

import (
	"fmt"
	"os"
	"time"
	"work_report/config"
	_ "work_report/docs"
	"work_report/middlewares/exception"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(exception.Recover())
	//访问日志吸入文件样式
	fileDir := config.LogPath + string(os.PathSeparator) + config.AppName + "_gin_access_" + time.Now().Format("2006-01-02") + ".log"
	file, _ := os.Create(fileDir)
	c := gin.LoggerConfig{
		Output:    file,
		SkipPaths: []string{"/test"},
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				params.ClientIP,
				params.TimeStamp.Format(time.RFC1123),
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		},
	}
	router.Use(gin.LoggerWithConfig(c))
	//各种工具以及需要验签的部分
	idx := router.Group("/")
	{
		ENVIR := os.Getenv("ACTIVE")
		if ENVIR != "pro" || ENVIR != "uat" {
			//swagger-doc路由
			idx.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		idx.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"title":  "work_report首页",
				"env":    os.Getenv("ACTIVE"),
				"apollo": config.GetApolloString("ENVIRONMENT", ""),
			})
		})
	}
	wrProjectsRouter(router)
	wrUsersRouter(router)
	wrWorksRouter(router)
	wrAttendanceRouter(router)
	wrDayoffRouter(router)
	//增加页面模版
	//router.LoadHTMLGlob("views/**/*")
	return router
}
