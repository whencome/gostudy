package producer

import (
	"demo050/rabbitmq/def"

	"github.com/streadway/amqp"
)

// PublishProducer 适用于pub/sub模式下的消息发送
type PublishProducer struct {
	channel *amqp.Channel
	conf    *def.Config
}

func NewPublishProducer(channel *amqp.Channel, conf *def.Config) *PublishProducer {
	return &PublishProducer{
		channel: channel,
		conf:    conf,
	}
}

// Publish 发布消息
func (p *PublishProducer) Publish(message []byte) error {
	// 尝试创建交换机
	err := p.channel.ExchangeDeclare(
		p.conf.QueueName,
		// 交换机类型
		"fanout",
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
	// 调用channel 发送消息到队列中
	return p.channel.Publish(
		p.conf.Exchange,
		"",
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
