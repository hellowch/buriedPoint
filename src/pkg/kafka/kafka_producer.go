package main

import (
	"buriedPoint/src/constant"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

//生产者
func main()  {
	producer_test()
}

func producer_test()  {
	log.Println("producer_test")
	config := sarama.NewConfig()   //实例化个sarama的Config
	config.Producer.RequiredAcks = sarama.WaitForAll   //消息机制，0 1 -1 性能递减，数据健壮性递增
	config.Producer.Partitioner = sarama.NewRandomPartitioner  //随机分区器
	config.Producer.Return.Successes = true  //是否开启消息发送成功后通知 successes channel
	config.Producer.Return.Errors = true  //是否开启消息发送失败后通知 errors channel
	config.Version = sarama.V2_7_0_0   //kafka版本

	producer, err := sarama.NewAsyncProducer([]string{constant.KafKaUrl}, config)
	if err != nil {
		log.Println("producer_test create producer error :", err.Error())
		return
	}

	defer producer.AsyncClose()
	
	msg := &sarama.ProducerMessage{
		Topic: constant.KafkaTopic,
		Key: sarama.StringEncoder("go_test"),
	}
	
	value := "this is message"
	for {
		fmt.Scan(&value)
		msg.Value = sarama.ByteEncoder(value)
		
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			log.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			log.Printf("err: %s\n", fail.Err.Error())
		}
	}

}