package mongo

import (
	mongoPkg "buriedPoint/src/pkg/mongo"
	"fmt"
	"gopkg.in/mgo.v2"
)

func TestData(dbname string, tablename string, dataMap interface{}) error {
	c := mongoPkg.Mongo.DB(dbname).C(tablename)
	err := c.Insert(dataMap)
	return err
}


func FindData(mongo *mgo.Session, dbname string, tablename string) error {
	persons := []Person{}
	c := mongo.DB(dbname).C(tablename)
	err := c.Find(nil).All(&persons)
	if err != nil {
		return err
	}
	//fmt.Println(persons)
	for i,v := range persons {
		fmt.Println(i,"::",v.Name)
	}
	return nil
}

type Person struct {
	Name      string
	Phone     string
	City      string
	Age       int8
	IsMan     bool
	Interests []string
}