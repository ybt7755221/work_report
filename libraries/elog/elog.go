package elog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
	"work_report/libraries/efile"
	"work_report/libraries/wmail"

	"github.com/gin-gonic/gin"
)

type Elog struct {
	ReqHost   string      `json:"request_host"`
	ReqUrl    string      `json:"request_url"`
	ReqIp     string      `json:"request_ip"`
	ReqTime   string      `json:"request_time"`
	ReqMethod string      `json:"request_method"`
	ReqHeader http.Header `json:"request_header"`
	ErrMsg    string      `json:"error_msg"`
	FileMsg   []FileMsg   `json:"file_msg"`
}

type FileMsg struct {
	FileLevel int    `json:"file_level"`
	File      string `json:"file"`
	Line      int    `json:"line"`
	FuncName  string `json:"function_name"`
}

/**
 * 获取当前文件信息
 */
func GetFileInfo(skip int) FileMsg {
	pc, file, line, ok := runtime.Caller(skip)
	funcName := runtime.FuncForPC(pc).Name()
	errorFile := FileMsg{}
	if ok {
		errorFile.FileLevel = skip
		errorFile.File = file
		errorFile.Line = line
		errorFile.FuncName = funcName
	}
	return errorFile
}

func fileName() string {
	return efile.LogFileName("errors")
}

/**
 * 创建错误日志，写入文件
 */
func New(errMsg string, errFile FileMsg) {
	elogStruct := Elog{}
	elogStruct.FileMsg[0] = errFile
	go err(errMsg, elogStruct)
}

func Newf(format string, errMsg ...interface{}) {
	msg := fmt.Sprintf(format, errMsg...)
	New(msg, FileMsg{})
}

func ErrMail(errMsg string, elogStruct Elog) {
	go func() {
		err(errMsg, elogStruct)
		elogStruct.ErrMsg = errMsg
		//errStr := fmt.Sprintf("Request URI : %s \n", elogStruct.ReqUrl)
		//errStr += fmt.Sprintf("Request Method : %s \n", elogStruct.ReqMethod)
		//errStr += fmt.Sprintf("Request Header : %s \n", elogStruct.ReqHeader)
		//errStr += fmt.Sprintf("Request IP : %s \n", elogStruct.ReqIp)
		//errStr += fmt.Sprintf("Request Time : %s \n", elogStruct.ReqTime)
		//errStr += fmt.Sprintf("error file : %s  Line: %d \n", elogStruct.File, elogStruct.Line)
		//errStr += fmt.Sprintf("Request FuncName : %s \n", elogStruct.FuncName)
		//errStr += fmt.Sprintf("Request URI : %s", elogStruct.ErrMsg)
		wmail.SendErrMail(elogStruct)
	}()
}

func err(errMsg string, elogStruct Elog) {
	var errByte []byte
	elogStruct.ErrMsg = errMsg
	errByte, _ = json.Marshal(elogStruct)
	//写入log文件
	_ = efile.WriteFile(fileName(), errByte, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
}

/**
 * 获取所有信息，http和file
 */
func GetAllInfo(c *gin.Context) Elog {
	errStruct := GetHttpInfo(c)
	errFileMsg := make([]FileMsg, 0)
	for i := 0; i < 6; i++ {
		fileInfo := GetFileInfo(i)
		errFileMsg = append(errFileMsg, fileInfo)
	}
	errStruct.FileMsg = errFileMsg
	return errStruct
}

/**
 * 获取网络信息
 */
func GetHttpInfo(c *gin.Context) Elog {
	msgStruct := new(Elog)
	msgStruct.ReqHost = c.Request.Host
	msgStruct.ReqUrl = c.Request.RequestURI
	msgStruct.ReqMethod = c.Request.Method
	msgStruct.ReqHeader = c.Request.Header
	msgStruct.ReqTime = time.Now().Format("2006-01-02 15:04:05")
	return *msgStruct
}
