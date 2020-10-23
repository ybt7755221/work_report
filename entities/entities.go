package entities

import "sync"

// swagger:response ApiResponse
type ApiResonse struct {
	Code int         `json:"code" example:"success：1000"`
	Msg  string      `json:"msg" example:"请求成功/失败"`
	Data interface{} `json:"data"`
}

type Pagination struct {
	PageNum  int
	PageSize int
	SortStr  string
}

const (
	EntityIsOk              = 1000
	EntityParametersMissing = 1001
	EntityTokenMissing      = 1002
	EntitySystemError       = 1003
	EntityPanic             = 1004
	EntityUnauthorized      = 1401
	EntityForbidden         = 1403
	EntityTimeout           = 1504
	EntityFailure           = 1100
)

var lang string
var once sync.Once

func init() {
	once.Do(func() {
		lang = "cn"
	})
}

func GetStatusMsg(code int) string {
	return statusMsg[lang][code]
}

var statusMsg = map[string]map[int]string{
	"cn": {
		EntityIsOk:              "请求成功",
		EntityParametersMissing: "缺少请求参数",
		EntityUnauthorized:      "签名验证错误",
		EntityForbidden:         "请求被禁止",
		EntityTimeout:           "请求超时",
		EntityTokenMissing:      "缺少token值",
		EntityFailure:           "请求失败",
		EntitySystemError:       "系统错误",
		EntityPanic:             "Panic报错",
	},
	"en": {
		EntityIsOk:              "Request Success",
		EntityParametersMissing: "The Some Parameters is Missing",
		EntityUnauthorized:      "Request Unauthorized",
		EntityForbidden:         "Request Forbidden",
		EntityTimeout:           "Request Timeout",
		EntityTokenMissing:      "The token is missing",
		EntityFailure:           "Request Failure",
		EntitySystemError:       "System Error",
		EntityPanic:             "Panic Error",
	},
}
