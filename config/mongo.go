package config

const (
	Log       = "sys_log"
	SYSTEMLOG = "system_log"
)

type MgoStruct struct {
	Name           string
	Host           string
	Port           string
	User           string
	Pass           string
	PoolLimit      string
	Timeout        string
	Direct         string
	ReplicaSetName string
	Group          string
}

var GMConfig map[string]MgoStruct

func init() {
	GMConfig = make(map[string]MgoStruct, 0)
	GMConfig[Log] = MgoStruct{
		Name:           GetApolloString("MONGO_LOG_NAME", ""),
		Host:           GetApolloString("MONGO_LOG_HOST", ""),
		Port:           GetApolloString("MONGO_LOG_PORT", ""),
		User:           GetApolloString("MONGO_LOG_USER", ""),
		Pass:           GetApolloString("MONGO_LOG_PASS", ""),
		PoolLimit:      GetApolloString("MONGO_POOL_LIMIT", ""),
		Timeout:        GetApolloString("MONGO_TIMEOUT", ""),
		Direct:         GetApolloString("MONGO_DIRECT", ""),
		ReplicaSetName: GetApolloString("MONGO_LOG_REPLICASET", ""),
		Group:          GetApolloString("MONGO_LOG_HOST_PORT_GROUP", ""),
	}
	GMConfig[SYSTEMLOG] = MgoStruct{
		Name:           GetApolloString("MONGO_LOG_NAME", SYSTEMLOG),
		Host:           GetApolloString("MONGO_LOG_HOST", "localhost"),
		Port:           GetApolloString("MONGO_LOG_PORT", "27017"),
		User:           GetApolloString("MONGO_LOG_USER", "root"),
		Pass:           GetApolloString("MONGO_LOG_PASS", "123456"),
		PoolLimit:      GetApolloString("MONGO_POOL_LIMIT", ""),
		Timeout:        GetApolloString("MONGO_TIMEOUT", ""),
		Direct:         GetApolloString("MONGO_DIRECT", ""),
		ReplicaSetName: GetApolloString("MONGO_LOG_REPLICASET", ""),
		Group:          GetApolloString("MONGO_LOG_HOST_PORT_GROUP", ""),
	}
}
