package producer

import (
	"demo050/rabbitmq/def"

	"github.com/streadway/amqp"
)

// SimpleProducer 用于Simple以及Work模式下发送消息
type SimpleProducer struct {
	channel *amqp.Channel
	conf    *def.Config
}

func NewSimpleProducer(channel *amqp.Channel, conf *def.Config) *SimpleProducer {
	return &SimpleProducer{
		channel: channel,
		conf:    conf,
	}
}

// Publish 发布消息
func (p *SimpleProducer) Publish(message []byte) error {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := p.channel.QueueDeclare(
		p.conf.QueueName,
		//是否持久化
		p.conf.Durable,
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
	//调用channel 发送消息到队列中
	return p.channel.Publish(
		p.conf.Exchange,
		p.conf.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}
