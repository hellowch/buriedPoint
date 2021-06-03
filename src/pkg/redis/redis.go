package redis

import (
	"buriedPoint/src/constant"
	"github.com/go-redis/redis"
)

var RedisConn *redis.ClusterClient

func InitRedis()  {
	RedisConn = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        []string{constant.RedisUrl},
		Password:     constant.RedisPassword,
		PoolSize:     constant.RedisPoolSize,
		MinIdleConns: constant.RedisMinIdLeConns,
	})
}
