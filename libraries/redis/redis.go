package redis

import (
	"fmt"
	"work_report/config"

	"github.com/go-redis/redis"
)

var Cache *redis.Client
var Db *redis.Client

func init() {
	ReConnect()
}

func ReConnect() {
	if Cache == nil {
		Cache = connect("cache", 0)
	}
	if Db == nil {
		Db = connect("db", 1)
	}
}

func connect(key string, isDb int) *redis.Client {
	redConf := config.RedisConfMap[key]
	addr := redConf.Host + ":" + redConf.Port
	fmt.Println("Redis Addr: " + addr)
	redConn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",   // no password set
		DB:       isDb, // use default DB
	})
	return redConn
}
