package entities

type WrWorks struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId    int    `json:"user_id" xorm:"not null comment('用户id') index(idx_user_created) INT(10)"`
	ProjectId int    `json:"project_id" xorm:"not null comment('项目id') index(idx_project_created) INT(11)"`
	Title     string `json:"title" xorm:"not null default '' comment('工作title') VARCHAR(128)"`
	Url       string `json:"url" xorm:"not null default '' comment('url地址') VARCHAR(255)"`
	Progress  int    `json:"progress" xorm:"not null default 0 comment('工作进度') TINYINT(3)"`
	WorkType  int    `json:"work_type"`
	Backup    string `json:"backup" xorm:"not null default '' comment('备注') VARCHAR(255)"`
	Created   string `json:"created"`
	Updated   string `json:"updated" xorm:"updated"`
}

type WrWorksPageDao struct {
	List     []WrWorks `json:"list"`
	PageNum  int       `json:"page_num"`
	PageSize int       `json:"page_size"`
	Total    int64     `json:"total"`
}
