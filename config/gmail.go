package config

import (
	"fmt"
	"os"
	"strings"
)

type EmailConf struct {
	Host     string
	Port     int
	User     string
	Passwd   string
	To       string
	ErrTopic string
}

var EmailConfStruct EmailConf

func init() {
	var topic string
	if strings.ToLower(os.Getenv("ACTIVE")) == "pro" {
		topic = "正式"
	} else {
		topic = "测试"
	}
	EmailConfStruct = EmailConf{
		Host:     GetApolloString("MAIL_ALERT_HOST", ""),
		Port:     GetApolloInt("MAIL_ALERT_PORT_SSL", 0),
		User:     GetApolloString("MAIL_ALERT_USER", ""),
		Passwd:   GetApolloString("MAIL_ALERT_PASS", ""),
		To:       GetApolloString("MAIL_ALERT_SEND", ""),
		ErrTopic: fmt.Sprintf("【work_report系统 | %s】错误", topic),
	}
}
