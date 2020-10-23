package dao

type WrProjectsDao struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	ProjectName string `json:"project_name" xorm:"not null default '' comment('项目名称') VARCHAR(128)"`
	Created     string `json:"created" xorm:"created"`
	Updated     string `json:"updated" xorm:"updated"`
}
