package verify

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"work_report/config"

	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context) (string, string) {
	var rawStr string
	methodStr := c.Request.Method
	if methodStr == "GET" {
		rawStr = getParams(c.Request.URL.Query())
	} else {
		rawStr = getParams(c.Request.PostForm)
	}
	rawStr = fmt.Sprintf("%s&%s", rawStr, config.GetApolloString(config.Secret, ""))
	return rawStr, GenerateMD5(rawStr, 32)
}

//处理参数
func getParams(data url.Values) string {
	delete(data, "token")
	if len(data) > 0 {
		keys := make([]string, 0)
		for key, _ := range data {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		var paramsStr string
		for _, val := range keys {
			if val != "token" {
				paramsStr += val + "&" + data.Get(val) + "&"
			}
		}
		return paramsStr[0 : len(paramsStr)-1]
	} else {
		return ""
	}
}

//获取md5
func GenerateMD5(raw string, size int) string {
	md5H := md5.New()
	md5H.Write([]byte(raw))
	token := hex.EncodeToString(md5H.Sum(nil))
	if size == 16 {
		return token[8:16]
	}
	return token
}
