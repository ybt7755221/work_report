package config

import (
	"strconv"
	"work_report/libraries/apolloCli"
)

const (
	AppName      = "work_report"
	Secret       = "Dl*sCKW7C{SfYiPtYX*O5/71vG9&sm?2U"
	HttpPort     = "8989"
	Duration     = "TIMEOUT"
	LogPath      = "/Users/Burt/Work/logs"
	WechatUrl    = "WECHAT"
	WechattoUser = "WECHAT_TOUSER"
	WechatSecret = "WECHAT_SECRET"
	WechatAppkey = "WECHAT_APPKEY"
	KafkaUrl     = "KFKURL"
	KafKaProt    = "KFKPORT"
)

type JaegerConf struct {
	Host string
	Port string
	Type string
}

func init() {
	apolloCli.OptionInit()
}

func GetApolloString(key string, defValue string) string {
	apoCli := apolloCli.GetApolloConfig()
	if apoCli[key] == nil {
		return defValue
	}
	return apoCli[key].(string)
}

func GetApolloInt(key string, defValue int) int {
	value := GetApolloString(key, "")
	if len(value) == 0 {
		return defValue
	}
	num, _ := strconv.Atoi(value)
	return num
}
