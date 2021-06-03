package main

import (
	"buriedPoint/api/routers"
	"buriedPoint/src/pkg/mongo"
	"buriedPoint/src/pkg/mysql"
	"buriedPoint/src/pkg/redis"
)

func main()  {
	routers.InitRouter()

	mongo.InitMongo()
	mysql.InitMysql()
	redis.InitRedis()
}