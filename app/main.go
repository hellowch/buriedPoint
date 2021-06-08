package main

import (
	"buriedPoint/api/routers"
	"buriedPoint/src/pkg/mysql"
	"buriedPoint/src/pkg/redis"
)

func main()  {
	//mongo.InitMongo()
	mysql.InitMysql()
	redis.InitRedis()

	routers.InitRouter()
}