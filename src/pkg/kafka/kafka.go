package kafka

import (
	"buriedPoint/src/constant"
	"buriedPoint/src/models/mongo"
	"encoding/json"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
	"time"
)

var Producer sarama.AsyncProducer
var Consumer *cluster.Consumer
var Topic []string

func InitKafka()  {
	kafkaProducer()
	kafkaConsumer()
	go kafkaConsumerCluster()
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
		log.Println("kafkaProducer： 连接错误 :", err.Error())
		return
	}

	//defer Producer.AsyncClose()
}

//主要用于获取topic等数据列表
func kafkaConsumer()  {
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_7_0_0
	brokers := []string{constant.KafKaUrl}
	client, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Println("kafkaConsumer： 连接错误 :", err.Error())
		return
	}
	//获取topic列表
	var topicList []string
	topicList, err = client.Topics()
	if err != nil {
		log.Println("topic获取错误")
	}
	//__consumer_offsets是用于保存消费者偏移量的
	for i:=0;i<len(topicList);i++ {
		if topicList[i] != "__consumer_offsets" {
			Topic = append(Topic, topicList[i])
		}
	}
	log.Println(Topic)
	//defer consumer.Close()
}

func kafkaConsumerCluster()  {
	brokers := []string{constant.KafKaUrl}
	//topics := []string{constant.KafkaTopic}

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.CommitInterval=1*time.Second
	config.Consumer.Offsets.Initial=sarama.OffsetNewest
	config.Group.Return.Notifications = true

	var err error
	//第二个参数是groupId
	Consumer, err = cluster.NewConsumer(brokers, "consumer-group1", Topic, config)
	if err != nil {
		panic(err)
	}
	defer Consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 接收错误
	go func() {
		for err = range Consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// 打印一些rebalance的信息
	go func() {
		for ntf := range Consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// 消费消息
	for {
		select {
		case msg, ok := <-Consumer.Messages():
			if ok {
				log.Printf("kafkaConsumerCluster： msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				BPInsertMongoData(msg.Value)
				Consumer.MarkOffset(msg, "")   // 提交offset
			}
		case <-signals:
			return
		}
	}
}

//向读取kafka埋点数据，写入mongo
func BPInsertMongoData(value []byte) {
	dataMap := make(map[string]string)
	//json转map
	err := json.Unmarshal(value, &dataMap)
	if err != nil {
		log.Println("BPInsertMongoData：json转map失败:", err)
		return
	}
	//写入mongo
	err = mongo.InsertMongo(dataMap)
	if err != nil {
		log.Println("BPInsertMongoData：写入mongo失败:", err)
		return
	}
}
