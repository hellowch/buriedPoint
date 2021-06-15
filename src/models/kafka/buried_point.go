package kafka

import (
	"buriedPoint/src/constant"
	kafkaPkg "buriedPoint/src/pkg/kafka"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

func KafkaProducer(value map[string]string)  {
	msg := &sarama.ProducerMessage{
		Topic: constant.KafkaTopic,
		Key: sarama.StringEncoder("go_test"),
	}
	//将map转成json
	value["time"] = time.Now().Format("2006-01-02 15:04:05")

	mjson, _ := json.Marshal(value)
	msg.Value = sarama.ByteEncoder(mjson)

	kafkaPkg.Producer.Input() <- msg

	select {
	case suc := <-kafkaPkg.Producer.Successes():
		log.Printf("Producer offset: %d,  timestamp: %s, value: %s", suc.Offset, suc.Timestamp.String(), suc.Value)
	case fail := <-kafkaPkg.Producer.Errors():
		log.Printf("err: %s\n", fail.Err.Error())
	}
}

//向读取kafka埋点数据，写入mongo
//func BPInsertMongoData(value []byte) {
//	dataMap := make(map[string]string)
//	//json转map
//	err := json.Unmarshal(value, &dataMap)
//	if err != nil {
//		log.Println("Umarshal failed:", err)
//		return
//	}
//	//写入mongo
//	err = mongo.InsertMongo(dataMap)
//	if err != nil {
//		log.Println("mongo failed:", err)
//		return
//	}
//}
