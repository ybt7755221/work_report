package entities

type WrProjects struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	ProjectName string `json:"project_name" xorm:"not null default '' comment('项目名称') VARCHAR(128)"`
	TestTime	string `json:"test_time"`
	PublishTime	string `json:"publish_time"`
	Created     string `json:"created" xorm:"created"`
	Updated     string `json:"updated" xorm:"updated"`
}

type WrProjectsPageDao struct {
	List     []WrProjects `json:"list"`
	PageNum  int          `json:"page_num"`
	PageSize int          `json:"page_size"`
	Total    int64        `json:"total"`
}
