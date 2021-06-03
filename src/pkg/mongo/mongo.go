package mongo

import (
	"buriedPoint/src/constant"
	"gopkg.in/mgo.v2"
	"log"
)


type Person struct {
	Name      string
	Phone     string
	City      string
	Age       int8
	IsMan     bool
	Interests []string
}

var Mongo *mgo.Session

func InitMongo() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{constant.MongoUrl},
		Timeout:  constant.MongoTimeout,
		Username: constant.MongoUsername,
		Password: constant.MongoPassword,
	}

	//连接
	Mongo, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Printf("CreateSession failed:%n", err)
	}

	//连接池限制
	Mongo.SetPoolLimit(constant.MongoMaxPoolSize)
	defer Mongo.Close()

	Mongo.SetMode(mgo.Monotonic, true)

	//调用测试
	err = createData()
	if err != nil {
		log.Println(err)
	}
}

func createData() error {

	persons := []Person{
		Person{Name: "Tony", Phone: "123432", City: "Shanghai", Age: 33, IsMan: true, Interests: []string{"music", "tea", "collection"}},
		Person{Name: "Mary", Phone: "232562", City: "Beijing", Age: 43, IsMan: false, Interests: []string{"sport", "film"}},
		Person{Name: "Tom", Phone: "123432", City: "Suzhou", Age: 22, IsMan: true, Interests: []string{"music", "reading"}},
		Person{Name: "Bob", Phone: "123432", City: "Hangzhou", Age: 32, IsMan: true, Interests: []string{"shopping", "coffee"}},
		Person{Name: "Alex", Phone: "15772", City: "Shanghai", Age: 21, IsMan: true, Interests: []string{"music", "chocolate"}},
		Person{Name: "Alice", Phone: "43456", City: "Shanghai", Age: 42, IsMan: false, Interests: []string{"outing", "tea"}},
		Person{Name: "Ingrid", Phone: "123432", City: "Shanghai", Age: 22, IsMan: false, Interests: []string{"travel", "tea"}},
		Person{Name: "Adle", Phone: "123432", City: "Shanghai", Age: 20, IsMan: false, Interests: []string{"game", "coffee", "sport"}},
		Person{Name: "Smith", Phone: "54223", City: "Fuzhou", Age: 54, IsMan: true, Interests: []string{"music", "reading"}},
		Person{Name: "Bruce", Phone: "123432", City: "Shanghai", Age: 31, IsMan: true, Interests: []string{"film", "tea", "game", "shoping", "reading"}},
	}

	cloneMongo := Mongo.Clone()
	c := cloneMongo.DB( "go_test").C( "go_test")

	for _, item := range persons {
		err := c.Insert(&item)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

