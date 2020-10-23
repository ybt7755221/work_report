package gutil

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//结构体赋值结构体
func BeanUtil(out interface{}, in interface{}) {
	outType := reflect.TypeOf(out).Elem()
	outValue := reflect.ValueOf(out).Elem()
	inType := reflect.TypeOf(in).Elem()
	inValue := reflect.ValueOf(in).Elem()
	outNum := outType.NumField()
	for i := 0; i < outNum; i++ {
		outFieldInfo := outType.Field(i)
		inTypeInfo, ok := inType.FieldByName(outFieldInfo.Name)
		if ok {
			outTypeString := outFieldInfo.Type.String()
			inTypeString := inTypeInfo.Type.String()
			inVal := inValue.FieldByName(outFieldInfo.Name)
			if outTypeString == inTypeString {
				outValue.FieldByName(outFieldInfo.Name).Set(reflect.Value(inValue.FieldByName(outFieldInfo.Name)))
			} else {
				var val interface{}
				switch outTypeString {
				case "int":
					if inTypeString == "string" {
						val, _ = strconv.Atoi(inVal.String())
					} else {
						val = int(inVal.Int())
					}
				case "int32":
					if inTypeString == "string" {
						val, _ = strconv.ParseInt(inVal.String(), 10, 32)
					} else {
						val = int32(inVal.Int())
					}
				case "int64":
					if inTypeString == "string" {
						val, _ = strconv.ParseInt(inVal.String(), 10, 64)
					} else {
						val = int64(inVal.Int())
					}
				case "string":
					if inTypeString == "time.Time" {
						val = inVal.Interface().(time.Time).Format("2006-01-02 15:04:05")
					} else {
						val = inVal.String()
					}
				case "float32":
					val, _ = strconv.ParseFloat(inVal.String(), 32)
				case "float64":
					val, _ = strconv.ParseFloat(inVal.String(), 64)
				case "time.Time":
					tmpValue := inVal.String()
					if len(tmpValue) == 10 && strings.Index(tmpValue, "-") == -1 {
						intTm, _ := strconv.ParseInt(tmpValue, 10, 64)
						tm := time.Unix(intTm, 0)
						tmpValue = tm.Format("2006-01-02 15:04:05")
					}
					val, _ = time.Parse("2006-01-02 15:04:05", tmpValue)
				default:
					val = nil
				}
				outValue.FieldByName(outFieldInfo.Name).Set(reflect.ValueOf(val))
			}
		}
	}
}

//
func TwoJson(out interface{}, in interface{}) {
	byte, _ := json.Marshal(in)
	json.Unmarshal(byte, out)
}

//首字母小写
func FirstToLower(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

//首字母大写
func FirstToUpper(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

//获取指定周第一天和最后一天
func GetWeekDay(n int) (string, string) {
	now := time.Now()
	switch n {
	case 1:
		now = time.Now().AddDate(0, 0, -7)
		break
	case 3:
		now = time.Now().AddDate(0, 0, 7)
		break
	default:
		break
	}
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		offset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

