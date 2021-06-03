package redis

import "github.com/go-redis/redis"

var RedisConn *redis.ClusterClient

func init()  {
	RedisConn = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        []string{"weichenhao.cn:6380"},
		Password:     "",
		PoolSize:     100,
		MinIdleConns: 50,
	})
}
