package config

type MysqlConf struct {
	Host    string
	Port    string
	Name    string
	User    string
	Passwd  string
	Charset string
	OpenMax int
	IdleMax int
}

const (
	Gin         = "work_reports"
	DefPageSize = 50
)

var MysqlConfMap map[string]MysqlConf

func init() {
	//库操作
	msqConfMap := map[string]MysqlConf{
		Gin: {
			Host:    GetApolloString("DB_HOST", "127.0.0.1"),
			Port:    GetApolloString("DB_PORT", "3357"),
			Name:    GetApolloString("DB_NAME", "work_reports"),
			User:    GetApolloString("DB_USER", "root"),
			Passwd:  GetApolloString("DB_PASS", "root"),
			Charset: "utf8",
			OpenMax: GetApolloInt("MYSQL_MAX_OPEN_CONN", 100),
			IdleMax: GetApolloInt("MYSQL_MAX_IDEL_CONN", 60),
		},
	}
	MysqlConfMap = msqConfMap
}
