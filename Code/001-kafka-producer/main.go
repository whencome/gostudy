package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	// 发送完数据需要leader和follow都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的消息将在success channel返回
	config.Producer.Return.Successes = true

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test_topic"
	msg.Value = sarama.StringEncoder("this is a test log from " + time.Now().Format("2006-01-02 15:04:05"))

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("connect to kafka failed: ", err)
		return
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed: ", err)
		return
	}
	fmt.Printf("pid = %v, offset = %v\n", pid, offset)

}
