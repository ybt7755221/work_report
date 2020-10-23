package controllers

import (
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	. "work_report/entities"
	"work_report/libraries/efile"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// swagger: success response
type SgrResp struct {
	Code int         `json:"code" example:"1000"`
	Msg  string      `json:"msg" example:"请求成功"`
	Data interface{} `json:"data"`
}

type accesslog struct {
	ReqUrl    string      `json:"request_url"`
	ReqIp     string      `json:"request_ip"`
	ReqTime   string      `json:"request_time"`
	ReqMethod string      `json:"request_method"`
	ReqHeader http.Header `json:"request_header"`
	ReqBody   url.Values  `json:"request_body"`
	ReqJson   interface{} `json:"request_body_json"`
	ResCode   int         `json:"response_code"`
	ResData   ApiResonse  `json:"response_data"`
}

func getParamsNew(c *gin.Context, dbEnt interface{}) {
	//遍历entities中的json
	var field string
	gType := reflect.TypeOf(dbEnt).Elem()
	gValue := reflect.ValueOf(dbEnt).Elem()
	structNum := gValue.NumField()
	for i := 0; i < structNum; i++ {
		field = gType.Field(i).Tag.Get("json")
		if c.Query(field) != "" {
			switch gType.Field(i).Type.String() {
			case "int":
				val, _ := strconv.Atoi(c.Query(field))
				gValue.Field(i).Set(reflect.ValueOf(val))
			case "int32":
				val, _ := strconv.ParseInt(c.Query(field), 10, 32)
				gValue.Field(i).Set(reflect.ValueOf(val))
			case "int64":
				val, _ := strconv.ParseInt(c.Query(field), 10, 64)
				gValue.Field(i).Set(reflect.ValueOf(val))
			case "string":
				gValue.Field(i).Set(reflect.ValueOf(c.Query(field)))
			case "float32":
				val, _ := strconv.ParseFloat(c.Query(field), 32)
				gValue.Field(i).Set(reflect.ValueOf(val))
			case "float64":
				val, _ := strconv.ParseFloat(c.Query(field), 64)
				gValue.Field(i).Set(reflect.ValueOf(val))
			case "time.Time":
				tmpValue := c.Query(field)
				if len(tmpValue) == 10 && strings.Index(tmpValue, "-") == -1 {
					intTm, _ := strconv.ParseInt(tmpValue, 10, 64)
					tm := time.Unix(intTm, 0)
					tmpValue = tm.Format("2006-01-02 15:04:05")
				}
				val, _ := time.Parse("2006-01-02 15:04:05", tmpValue)
				gValue.Field(i).Set(reflect.ValueOf(val))
			default:
			}
		}
	}
}

func getParams(c *gin.Context, fields []string, dbEnt interface{}) map[string]string {
	condition := make(map[string]string)
	if len(fields) > 0 {
		//field存在，以field为准
		for _, val := range fields {
			condition[val] = c.Query(val)
		}
	} else {
		//如果field为空，遍历entities中的json
		var field string
		gType := reflect.TypeOf(dbEnt)
		gValue := reflect.ValueOf(dbEnt)
		structNum := gValue.NumField()
		for i := 0; i < structNum; i++ {
			field = gType.Field(i).Tag.Get("json")
			condition[field] = c.Query(field)
		}
	}
	return condition
}

func resJson(c *gin.Context, httpCode int, data ApiResonse) {
	c.JSON(httpCode, data)
	go func() {
		writeReqLog(c, httpCode, data, true)
	}()
}

func resResult(c *gin.Context, code int, msg string, data interface{}) {
	dataStruct := ApiResonse{code, msg, data}
	resJson(c, http.StatusOK, dataStruct)
}

func resSuccess(c *gin.Context, data interface{}) {
	dateStruct := ApiResonse{EntityIsOk, GetStatusMsg(EntityIsOk), data}
	resJson(c, http.StatusOK, dateStruct)
}

func resError(c *gin.Context, code int, msg string) {
	dateStruct := ApiResonse{code, msg, gin.H{}}
	resJson(c, http.StatusOK, dateStruct)
}

/**
 * 字符格式转时间
 * "2006-01-02 15:04:05"
 */
func string2Time(tSter string) (theTime time.Time) {
	loc, _ := time.LoadLocation("Local")
	theTime, _ = time.ParseInLocation("2006-01-02", tSter, loc)
	return
}

//请求日志
func writeReqLog(c *gin.Context, code int, data ApiResonse, db bool) {
	al := accesslog{}
	al.ReqUrl = c.Request.URL.Path
	al.ReqBody = c.Request.PostForm
	al.ReqTime = time.Now().Format("2006-01-02")
	al.ReqIp = c.ClientIP()
	al.ReqHeader = c.Request.Header
	al.ReqMethod = c.Request.Method
	al.ResCode = code
	al.ResData = data
	if al.ReqBody == nil {
		c.ShouldBindBodyWith(&al.ReqJson, binding.JSON)
	}
	file := efile.LogFileName("requests")
	if db == true {
		recordMongo(al)
	}
	//写入log文件
	_ = efile.WriteFile(file, al, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
}

//记录mongo
func recordMongo(data accesslog) {
	//mgoName := "clue_api_request_log_" + time.Now().Format("200601")
	//mongo.Insert(config.Log, mgoName, data)
}

//获取update/create参数,通过反射，返回响应结构体
func getPostStructData(c *gin.Context, in interface{}) {
	postMap := make(map[string]interface{})
	gValue := reflect.ValueOf(in).Elem()
	num := gValue.NumField()
	for i := 0; i < num; i++ {
		fieldInfo := gValue.Type().Field(i)
		name := fieldInfo.Tag.Get("json")
		postMap[name] = c.PostForm(name)
	}
	//for key, value :=range c.Request.PostForm {
	//	postMap[key] = value[0]
	//}
	mapToStruct(in, postMap)
	return
}

//map赋值struct
func mapToStruct(ptr interface{}, fields map[string]interface{}) {
	gValue := reflect.ValueOf(ptr).Elem()
	num := gValue.NumField()
	for i := 0; i < num; i++ {
		fieldInfo := gValue.Type().Field(i)
		jName := fieldInfo.Tag.Get("json")
		if jName == "" {
			jName = strings.ToLower(fieldInfo.Name)
		}
		if value, ok := fields[jName]; ok {
			if reflect.TypeOf(value) == gValue.Field(i).Type() {
				gValue.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			} else {
				if gValue.Field(i).Type().String() == "string" {
					gValue.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value.(string)))
				}
				if gValue.Field(i).Type().String() == "int" {
					val, _ := strconv.Atoi(value.(string))
					gValue.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(val))
				}
				if gValue.Field(i).Type().String() == "int32" {
					val, _ := strconv.ParseInt(value.(string), 10, 32)
					gValue.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(val))
				}
				if gValue.Field(i).Type().String() == "int64" {
					val, _ := strconv.ParseInt(value.(string), 10, 64)
					gValue.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(val))
				}
				if gValue.Field(i).Type().String() == "float32" {
					val, _ := strconv.ParseFloat(value.(string), 32)
					gValue.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(val))
				}
				if gValue.Field(i).Type().String() == "float64" {
					val, _ := strconv.ParseFloat(value.(string), 64)
					gValue.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(val))
				}
				if gValue.Field(i).Type().String() == "time.Time" {
					tmpValue := value.(string)
					if len(tmpValue) == 10 && strings.Index(tmpValue, "-") == -1 {
						intTm, _ := strconv.ParseInt(tmpValue, 10, 64)
						tm := time.Unix(intTm, 0)
						tmpValue = tm.Format("2006-01-02 15:04:05")
					}
					theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", tmpValue, time.Local)
					gValue.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(theTime))
				}
			}
		}
	}
	return
}
