package config

type RedisConf struct {
	Addr        string
	MinIdle     int
	MaxActive   int
	IdleTimeout int
}
type RedisENV struct {
	Host string
	Port string
}

var RedisConfMap = make(map[string]RedisENV)
var RedisConfig *RedisConf

func init() {
	RedisConfig = &RedisConf{
		MinIdle:     5,
		MaxActive:   100,
		IdleTimeout: 600,
	}
	RedisConfMap["cache"] = RedisENV{
		Host: GetApolloString("REDIS_SERVER", "127.0.0.1"),
		Port: GetApolloString("REDIS_PORT", "6380"),
	}
	RedisConfMap["db"] = RedisENV{
		Host: GetApolloString("SREDIS_DB_SERVER", "127.0.0.1"),
		Port: GetApolloString("REDIS_DB_PORT", "6380"),
	}
}
