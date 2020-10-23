package dao

type WrUsersDao struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username string `json:"username" xorm:"not null default '' comment('用户名') VARCHAR(128)"`
	Mobile   string `json:"mobile" xorm:"not null default '13000000000' comment('手机号') VARCHAR(20)"`
	Email    string `json:"email" xorm:"not null default '' comment('邮箱地址') VARCHAR(128)"`
	Created  string `json:"created" xorm:"created"`
	Updated  string `json:"updated" xorm:"updated"`
}
