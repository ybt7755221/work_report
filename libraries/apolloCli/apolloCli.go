package apolloCli

import (
	"fmt"
	"os"
	"strings"

	"github.com/shima-park/agollo"
)

const (
	AppId         = "AppId"
	IpFAT         = "http://localhost"
	IpUAT         = "http://localhost"
	IpPRO         = "http://localhost"
	NameSpacename = "application"
	BackUpFile    = "/etc/application.agollo"
)

var confMap map[string]interface{}

func OptionInit() map[string]interface{} {
	env := os.Getenv("ACTIVE")
	var Ip string
	if strings.ToLower(env) == "pro" {
		Ip = IpPRO
	} else if strings.ToLower(env) == "uat" {
		Ip = IpUAT
	} else {
		Ip = IpFAT
	}
	apoCli, err := agollo.New(Ip, AppId,
		agollo.BackupFile(BackUpFile),
		agollo.FailTolerantOnBackupExists(),
		agollo.AutoFetchOnCacheMiss(),
	)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("apollo start is successd")
	}
	confMap = apoCli.GetNameSpace(NameSpacename)
	apoCli.Start() // Start后会启动goroutine监听变化，
	go func() {
		watchCh := apoCli.Watch()
		for {
			select {
			case resp := <-watchCh:
				txtFile := "【Apollo Update】Apollo has modified！Namespace is " + resp.Namespace
				fmt.Println(txtFile)
				confMap = apoCli.GetNameSpace(NameSpacename)
			}
		}
	}()
	return confMap
}

func GetApolloConfig() map[string]interface{} {
	return confMap
}
