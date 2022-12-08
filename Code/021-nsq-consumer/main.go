package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

var (
	tcpNsqdAddr = "127.0.0.1:4150"
	topic       = "test_topic"
	exitChan    = make(chan struct{})
)

type MsgHandler struct {
	Name string
}

// HandleMessage 处理消息
func (h *MsgHandler) HandleMessage(message *nsq.Message) error {
	fmt.Printf("%s: msg.Timestamp=%v, msg.nsqaddress=%s,msg.body=%s \n", h.Name, time.Unix(0, message.Timestamp).Format("2006-01-02 03:04:05"), message.NSQDAddress, string(message.Body))
	return nil
}

func main() {
	// 初始化配置
	config := nsq.NewConfig()
	// 创建消费者
	consumer, err := nsq.NewConsumer(topic, "test_channel", config)
	if err != nil {
		fmt.Printf("create new consumer failed: %s\n", err)
		return
	}

	// 添加消息处理器
	consumer.AddHandler(&MsgHandler{Name: "my_handler"})
	// 连接nsqd
	err = consumer.ConnectToNSQD(tcpNsqdAddr)
	if err != nil {
		fmt.Printf("connect to %s failed: %s\n", tcpNsqdAddr, err)
		return
	}
	<-exitChan
}
