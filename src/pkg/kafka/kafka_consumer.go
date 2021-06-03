package main

import (
	"buriedPoint/src/constant"
	"github.com/Shopify/sarama"
	"log"
)

func main()  {
	consumer_test()
}

func consumer_test()  {
	log.Println("consumer_test")
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_7_0_0

	consumer, err := sarama.NewConsumer([]string{constant.KafKaUrl}, config)
	if err != nil {
		log.Println("consumer_test create consumer error :", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition(constant.KafkaTopic,0,sarama.OffsetOldest)
	if err != nil {
		log.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			log.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			log.Printf("err :%s\n", err.Error())
		}
	}
}