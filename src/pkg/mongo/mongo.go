package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type Student struct {
	Name   string   `json:"name"`
	Age    int		`json:"age"`
	Sid    string	`json:"sid"`
	Status int		`json:"status"`
}

func main()  {
	mongo, err := mgo.Dial("mongodb://admin:123456@weichenhao.cn:27017")

	defer mongo.Close()
	if err != nil {
		fmt.Println("1:",err)
		return
	}
	client := mongo.DB("go_test").C("go_test")

	data := Student{
		Name: "wch",
		Age: 21,
		Sid: "12345667",
		Status: 1,
	}

	err = client.Insert(&data)
	if err != nil {
		fmt.Println("2:",err)
		return
	}
}
