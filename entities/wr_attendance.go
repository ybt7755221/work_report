package entities

type WrAttendance struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId   int    `json:"user_id" xorm:"not null index INT(11)"`
	Overtime string `json:"overtime" xorm:"not null comment('加班时间') DATE"`
	Hours    int    `json:"hours" xorm:"not null default 0 comment('加班时长h') TINYINT(2)"`
	Status   int    `json:"status" xorm:"not null default 1 comment('状态 1未使用； 0已用完；2 部分使用') TINYINT(1)"`
	Created  string `json:"created" xorm:"created"`
	Updated  string `json:"updated" xorm:"updated"`
}

type WrAttendancePageDao struct {
	List     []WrAttendance `json:"list"`
	PageNum  int            `json:"page_num"`
	PageSize int            `json:"page_size"`
	Total    int64          `json:"total"`
}
