package constant

import "time"

const (
	Url = "weichenhao.cn:"

	KafKaUrl = Url + "9092"
	KafkaTopic = "kafka_go_test"

	MongoUrl = Url + "27017"
	MongoTimeout = 60 * time.Second
	MongoUsername = "root"
	MongoPassword = "123456"
	MongoMaxPoolSize = 300

	MysqlUrl = "root:123456@tcp(" + Url + "3310" + ")/chargeMs?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlMaxIdleConns = 10
	MysqlMaxOpenConns = 100
	MysqlConnMaxLifetime = time.Hour

	RedisUrl = Url + "6380"
	RedisPassword = ""
	RedisPoolSize = 100
	RedisMinIdLeConns = 50

)