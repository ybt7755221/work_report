package redis

import (
	redisServ "github.com/go-redis/redis"
)

type IRedisGeo interface {
	GeoAdd(string, ...*redisServ.GeoLocation) (int64, error)
	GeoPos(string, ...string) ([]*redisServ.GeoPos, error)
	GeoHash(string, ...string) ([]string, error)
	GeoDist(string, string, string, string) (float64, error)
	GeoRadius(string, float64, float64) ([]redisServ.GeoLocation, error)
	GeoRadiusByMember(string, string) ([]redisServ.GeoLocation, error)
}

type RedisGeo struct {
	IsDB        bool
	RadiusQuery *redisServ.GeoRadiusQuery
}

func (s *RedisGeo) client() *redisServ.Client {
	if s.IsDB {
		return Db
	}
	return Cache
}
func (s *RedisGeo) GeoDist(key string, name1 string, name2 string, unit string) (float64, error) {
	return s.client().GeoDist(key, name1, name2, unit).Result()
}
func (s *RedisGeo) GeoHash(key string, name ...string) ([]string, error) {
	return s.client().GeoHash(key, name...).Result()
}

func (s *RedisGeo) GeoPos(key string, name ...string) ([]*redisServ.GeoPos, error) {
	return s.client().GeoPos(key, name...).Result()
}

func (s *RedisGeo) GeoAdd(key string, getLocation ...*redisServ.GeoLocation) (int64, error) {
	return s.client().GeoAdd(key, getLocation...).Result()
}

func (s *RedisGeo) GeoRadius(key string, lat float64, lng float64) ([]redisServ.GeoLocation, error) {
	return s.client().GeoRadius(key, lat, lng, s.RadiusQuery).Result()
}

func (s *RedisGeo) GeoRadiusByMember(key string, member string) ([]redisServ.GeoLocation, error) {
	return s.client().GeoRadiusByMember(key, member, s.RadiusQuery).Result()
}
