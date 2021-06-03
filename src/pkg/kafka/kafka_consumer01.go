package main

import (
	"buriedPoint/src/constant"
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
	"time"
)

//消费者组
func main()  {
	consumer_test01()
}

func consumer_test01()  {
	fmt.Println("consumer_test01")
	brokers := []string{constant.KafKaUrl}
	topics := []string{constant.KafkaTopic}

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.CommitInterval=1*time.Second
	config.Consumer.Offsets.Initial=sarama.OffsetNewest
	config.Group.Return.Notifications = true

	//第二个参数是groupId
	consumer, err := cluster.NewConsumer(brokers, "consumer-group1", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 接收错误
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// 打印一些rebalance的信息
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// 消费消息
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				log.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				consumer.MarkOffset(msg, "")   // 提交offset
			}
		case <-signals:
			return
		}
	}

}
