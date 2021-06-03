package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

const (
	MongoDBHosts = "weichenhao.cn:27017"
	AuthUserName = "admin"
	AuthPassword = "123456"
	MaxPoolSize  = 300
)

type Person struct {
	Name      string
	Phone     string
	City      string
	Age       int8
	IsMan     bool
	Interests []string
}

func main() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	mongo, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession failed:%n", err)
	}

	mongo.SetPoolLimit(MaxPoolSize)
	defer mongo.Close()

	mongo.SetMode(mgo.Monotonic, true)

	err = createData(mongo, "go_test", "go_test")
	if err != nil {
		log.Fatal(err)
	}
}

func createData(mongo *mgo.Session, dbname string, tablename string) error {

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

	cloneMongo := mongo.Clone()
	c := cloneMongo.DB(dbname).C(tablename)

	for _, item := range persons {
		err := c.Insert(&item)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

