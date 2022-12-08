package main

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

var (
	tcpNsqdAddr = "127.0.0.1:4150"
	topic       = "test_topic"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(tcpNsqdAddr, config)
	if err != nil {
		fmt.Printf("create producer failed: %s \n", err)
		return
	}

	msgData := "this is a test message"
	err = producer.Publish(topic, []byte(msgData))
	if err != nil {
		fmt.Printf("produce message failed: %s \n", err)
		return
	}
	fmt.Println("produce message success")
}
