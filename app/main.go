package main

import (
	"buriedPoint/api/routers"
	"buriedPoint/src/pkg/mongo"
	"buriedPoint/src/pkg/mysql"
	"buriedPoint/src/pkg/redis"
)

func main()  {
	mysql.InitMysql()
	redis.InitRedis()
	mongo.InitMongo()

	routers.InitRouter()
}