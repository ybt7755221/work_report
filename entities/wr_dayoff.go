package entities

type WrDayoff struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	AttendanceId int    `json:"attendance_id" xorm:"not null comment('加班id') index INT(11)"`
	Dayoff       string `json:"dayoff" xorm:"not null comment('休假开始时间') DATE"`
	Hours        int    `json:"hours" xorm:"not null default 0 comment('休假时间') TINYINT(2)"`
	Backup       string `json:"backup" xorm:"not null default '' comment('备注，休假原因') VARCHAR(255)"`
	Created      string `json:"created" xorm:"created"`
	Updated      string `json:"updated" xorm:"updated"`
}

type WrDayoffPageDao struct {
	List     []WrDayoff `json:"list"`
	PageNum  int        `json:"page_num"`
	PageSize int        `json:"page_size"`
	Total    int64      `json:"total"`
}
