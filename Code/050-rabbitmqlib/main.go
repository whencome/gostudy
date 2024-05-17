package main

import (
	"demo050/rabbitmq"
	"demo050/rabbitmq/def"
	"fmt"
	"time"
)

func main() {
	conf := &def.Config{
		MQUrl:     "amqp://root:root@127.0.0.1:5672/",
		QueueName: "simple_queue",
		Exchange:  "",
		Key:       "",
		Durable:   true,
		AutoAck:   false,
	}
	mq := rabbitmq.New(conf)
	if err := mq.Open(); err != nil {
		fmt.Printf("connect to mq fail: %v\n", err)
		return
	}
	go mq.Consume(func(msg []byte) error {
		fmt.Println("<<<<< ", time.Now().Format("2006-01-02 15:04:05"), " [RECV] : ", string(msg))
		return nil
	})
	// 发布10调消息
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("message %d at %s", i, time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(">>>>> SEND: ", msg)
		mq.Publish([]byte(msg))
		time.Sleep(time.Second * 1)
	}
	time.Sleep(time.Second * 3)
	mq.Close()
}
