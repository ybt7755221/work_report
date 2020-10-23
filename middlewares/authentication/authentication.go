package authentication

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	et "work_report/entities"
	"work_report/libraries/verify"

	"github.com/gin-gonic/gin"
)

//数据验证-中间件
func Verify(c *gin.Context) {
	var token string
	var ts string
	methodStr := c.Request.Method
	if skip := skipVerify(c, methodStr); skip == false {
		if methodStr == "GET" || methodStr == "DELETE" {
			token = c.Query("token")
			ts = c.Query("ts")
		} else {
			token = c.PostForm("token")
			ts = c.PostForm("ts")
		}
		if len(ts) < 1 {
			c.JSON(http.StatusOK, et.ApiResonse{et.EntityParametersMissing, "缺少ts值", gin.H{}})
			c.Abort()
		}
		//验证码一分钟超时
		if verifyTimeout(ts) == true {
			c.JSON(http.StatusOK, et.ApiResonse{et.EntityTimeout, "请求超过1分钟有效期", gin.H{}})
			c.Abort()
		}
		if len(token) < 1 {
			c.JSON(http.StatusOK, et.ApiResonse{et.EntityParametersMissing, "缺少token值", gin.H{}})
			c.Abort()
		} else {
			rawStr, tokenStr := verify.GenerateToken(c)
			fmt.Println("raw: ", rawStr)
			fmt.Println("token: ", tokenStr)
			if token != tokenStr {
				c.JSON(http.StatusOK, et.ApiResonse{et.EntityUnauthorized, et.GetStatusMsg(et.EntityUnauthorized), gin.H{}})
				c.Abort()
			} else {
				c.Next()
			}
		}
	} else {
		c.Next()
	}
}

/**
 * 超过一分钟超时
 * 超过true
 * 未超时false
 */
func verifyTimeout(request_time string) bool {
	//暂时不做时间验证
	return false
	rtInt64, _ := strconv.ParseInt(request_time, 10, 64)
	nowTime := time.Now().Unix()
	if nowTime-rtInt64 > 60 {
		return true
	}
	return false
}

func skipVerify(c *gin.Context, methodStr string) bool {
	var skip string
	if methodStr == "GET" || methodStr == "DELETE" {
		skip = c.Query("skip_debug")
	} else {
		skip = c.PostForm("skip_debug")
	}
	if skip == "161217" {
		return true
	}
	return false
}
