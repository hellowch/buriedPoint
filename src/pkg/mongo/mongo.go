//package main
package mongo

import (
	"buriedPoint/src/constant"
	"gopkg.in/mgo.v2"
	"log"
)


var Mongo *mgo.Session

func InitMongo() {
	var err error
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{constant.MongoUrl},
		Timeout:  constant.MongoTimeout,
		Username: constant.MongoUsername,
		Password: constant.MongoPassword,
	}
	//连接
	Mongo, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Printf("CreateSession failed:%n", err)
	}

	//连接池限制
	Mongo.SetPoolLimit(constant.MongoMaxPoolSize)
	Mongo.SetMode(mgo.Monotonic, true)
	//defer Mongo.Close()
}

