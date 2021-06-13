package kafka

import (
	"buriedPoint/src/constant"
	kafkaPkg "buriedPoint/src/pkg/kafka"
	"github.com/Shopify/sarama"
	"log"
)

func KafkaProducer(value string)  {
	msg := &sarama.ProducerMessage{
		Topic: constant.KafkaTopic,
		Key: sarama.StringEncoder("go_test"),
	}

	msg.Value = sarama.ByteEncoder(value)

	kafkaPkg.Producer.Input() <- msg

	select {
	case suc := <-kafkaPkg.Producer.Successes():
		log.Printf("offset: %d,  timestamp: %s, value: %s", suc.Offset, suc.Timestamp.String(), suc.Value)
	case fail := <-kafkaPkg.Producer.Errors():
		log.Printf("err: %s\n", fail.Err.Error())
	}
}

func KafkaConsumer()  {
	// 消费消息
	for {
		select {
		case msg, ok := <-kafkaPkg.Consumer.Messages():
			if ok {
				log.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				kafkaPkg.Consumer.MarkOffset(msg, "")   // 提交offset
			}
		//case <-signals:
		//	return
		}
	}
}