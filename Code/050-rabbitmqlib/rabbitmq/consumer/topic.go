package consumer

import (
	"demo050/rabbitmq/def"

	"github.com/streadway/amqp"
)

// TopicConsumer 适用于topic模式下的消息消费
type TopicConsumer struct {
	channel  *amqp.Channel
	conf     *def.Config
	exitChan chan bool
}

func NewTopicConsumer(channel *amqp.Channel, conf *def.Config) *TopicConsumer {
	return &TopicConsumer{
		channel: channel,
		conf:    conf,
	}
}

func (c *TopicConsumer) Consume(f def.ConsumeFunc) error {
	// 1.试探性创建交换机
	err := c.channel.ExchangeDeclare(
		c.conf.Exchange,
		//交换机类型
		"topic",
		c.conf.Durable,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// 2.试探性创建队列，这里注意队列名称不要写
	q, err := c.channel.QueueDeclare(
		"", //随机生产队列名称
		c.conf.Durable,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	//绑定队列到 exchange 中
	err = c.channel.QueueBind(
		q.Name,
		c.conf.Key,
		c.conf.Exchange,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	//消费消息
	msgs, err := c.channel.Consume(
		q.Name,
		"",
		c.conf.AutoAck,
		false,
		false,
		false,
		nil,
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

func (c *TopicConsumer) Stop() {
	close(c.exitChan)
}
