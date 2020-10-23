package exception

import (
	"encoding/json"
	"net/http"
	"os"
	et "work_report/entities"
	"work_report/libraries/efile"
	"work_report/libraries/elog"
	"work_report/libraries/wmail"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errStruct := elog.GetAllInfo(c)
				errByte, _ := json.Marshal(err)
				errMsg := string(errByte)
				errStruct.ErrMsg = errMsg
				//写入文件
				go func() {
					fileName := efile.LogFileName("painc")
					//写入log文件
					_ = efile.WriteFile(fileName, errStruct, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
					//发送邮件
					wmail.SendErrMail(errStruct)
				}()
				//页面输出
				c.JSON(http.StatusOK, et.ApiResonse{et.EntityPanic, et.GetStatusMsg(et.EntityPanic), errStruct})
				c.Abort()
			}
		}()
		c.Next()
	}
}
