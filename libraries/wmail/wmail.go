package wmail

import (
	"crypto/tls"
	"encoding/json"
	"strings"
	"work_report/config"

	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {
	conf := config.EmailConfStruct
	m := gomail.NewMessage()
	m.SetHeader("From", config.AppName+"<"+conf.User+">")
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文
	d := gomail.NewDialer(
		conf.Host,
		conf.Port,
		conf.User,
		conf.Passwd,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}

func SendErrMail(body interface{}) error {
	bodyByte, _ := json.Marshal(body)
	conf := config.EmailConfStruct
	mailTo := strings.Split(conf.To, ",")
	return SendMail(mailTo, conf.ErrTopic, string(bodyByte))
}
