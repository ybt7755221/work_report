package httpReq

import (
	"net/url"
	"sort"
	"strings"
	"time"
	"work_report/config"
	"work_report/libraries/mongo"
	"work_report/libraries/verify"

	"go.mongodb.org/mongo-driver/bson"
)

type WechatReq struct {
	Undata bool
	ToUser string
}

func (w *WechatReq) wechatSN(data map[string]string) string {
	var str string
	//将map key写入slice，对key排序
	keySli := make([]string, 0)
	for key, _ := range data {
		keySli = append(keySli, key)
	}
	sort.Strings(keySli)
	for _, key := range keySli {
		str += "&" + key + "=" + data[key]
	}
	str = str[1:] + config.GetApolloString(config.WechatSecret, "")
	if !w.Undata {
		str += time.Now().Format("2006-01-02")
	}
	return strings.ToLower(verify.GenerateMD5(str, 32))
}

func (w *WechatReq) SendWechat(data string) (map[string]interface{}, error) {
	touser := config.GetApolloString(config.WechattoUser, "")
	if len(w.ToUser) > 0 {
		touser = w.ToUser
	}
	params := map[string]string{
		"touser":  touser,
		"content": data,
		"appkey":  config.GetApolloString(config.WechatAppkey, ""),
	}
	sn := w.wechatSN(params)
	params["sn"] = sn
	urlValue := url.Values{}
	for key, value := range params {
		urlValue.Add(key, value)
	}
	//发送post请求
	var ireq IRequest
	var hreq Request
	w.Undata = false
	hreq.ContentType = FORM
	ireq = hreq
	httpUrl := config.GetApolloString(config.WechatUrl, "")
	res, err := ireq.PostForm(httpUrl, urlValue)
	//记录mongo
	go func() {
		mgoName := "wechat_log_" + time.Now().Format("200601")
		mongo.InsertOne(config.SYSTEMLOG, mgoName, bson.M{
			"url":           httpUrl,
			"request_data":  urlValue,
			"response_data": res,
			"from":          config.AppName,
			"request_time":  time.Now().Format("2006-01-02 15:04:05"),
		})
	}()
	return res, err
}
