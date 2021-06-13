package kafka

import (
	"buriedPoint/src/constant"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
	"time"
)

var Producer sarama.AsyncProducer
var Consumer *cluster.Consumer

func InitKafka()  {
	kafkaProducer()
	kafkaConsumer()
}

func kafkaProducer()  {
	//生产者配置
	config := sarama.NewConfig()   //实例化个sarama的Config
	config.Producer.RequiredAcks = sarama.WaitForAll   //消息机制，0 1 -1 性能递减，数据健壮性递增
	config.Producer.Partitioner = sarama.NewRandomPartitioner  //随机分区器
	config.Producer.Return.Successes = true  //是否开启消息发送成功后通知 successes channel
	config.Producer.Return.Errors = true  //是否开启消息发送失败后通知 errors channel
	config.Version = sarama.V2_7_0_0   //kafka版本

	var err error
	//生产者连接
	Producer, err = sarama.NewAsyncProducer([]string{constant.KafKaUrl}, config)
	if err != nil {
		log.Println("producer_test create producer error :", err.Error())
		return
	}

	//defer Producer.AsyncClose()
}

func kafkaConsumer()  {
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