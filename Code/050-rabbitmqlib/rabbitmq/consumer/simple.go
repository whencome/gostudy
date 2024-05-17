package consumer

import (
	"demo050/rabbitmq/def"

	"github.com/streadway/amqp"
)

// SimpleConsumer 适用于Simple以及Work模式下的消息消费
type SimpleConsumer struct {
	channel  *amqp.Channel
	conf     *def.Config
	exitChan chan bool
}

func NewSimpleConsumer(channel *amqp.Channel, conf *def.Config) *SimpleConsumer {
	return &SimpleConsumer{
		channel: channel,
		conf:    conf,
	}
}

func (c *SimpleConsumer) Consume(f def.ConsumeFunc) error {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := c.channel.QueueDeclare(
		c.conf.QueueName,
		//是否持久化
		c.conf.Durable,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		return err
	}

	//接收消息
	msgs, err := c.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		c.conf.AutoAck, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	//启用协程处理消息
	go func() {
		for d := range msgs {
			var err error = nil
			if f != nil {
				err = f(d.Body)
			}
			if !c.conf.AutoAck && err == nil {
				// 手动确认消息
				d.Ack(true)
			}
		}
	}()

	// wait for exit
	<-c.exitChan
	return nil
}

func (c *SimpleConsumer) Stop() {
	close(c.exitChan)
}
