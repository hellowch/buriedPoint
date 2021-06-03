package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

//生产者
func main()  {
	producer_test()
}

func producer_test()  {
	fmt.Println("producer_test")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_7_0_0

	producer, err := sarama.NewAsyncProducer([]string{"weichenhao.cn:9092"}, config)
	if err != nil {
		fmt.Println("producer_test create producer error :", err.Error())
		return
	}

	defer producer.AsyncClose()
	
	msg := &sarama.ProducerMessage{
		Topic: "kafka_go_test",
		Key: sarama.StringEncoder("go_test"),
	}
	
	value := "this is message"
	for {
		fmt.Scan(&value)
		msg.Value = sarama.ByteEncoder(value)
		
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}

}