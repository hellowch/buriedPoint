package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main()  {
	consumer_test()
}

func consumer_test()  {
	fmt.Println("consumer_test")
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_7_0_0

	consumer, err := sarama.NewConsumer([]string{"weichenhao.cn:9092"}, config)
	if err != nil {
		fmt.Println("consumer_test create consumer error :", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition("kafka_go_test",0,sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}