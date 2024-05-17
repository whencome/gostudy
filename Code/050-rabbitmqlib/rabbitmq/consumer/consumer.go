package consumer

import (
	"demo050/rabbitmq/def"

	"github.com/streadway/amqp"
)

type Consumer interface {
	// 消费消息
	Consume(f def.ConsumeFunc) error
	// 停止消费
	Stop()
}

// New 创建一个新的producer
func New(channel *amqp.Channel, conf *def.Config) Consumer {
	switch conf.Mode {
	case def.ModePublish:
		return NewSubscribeConsumer(channel, conf)
	case def.ModeRouting:
		return NewRoutingConsumer(channel, conf)
	case def.ModeTopic:
		return NewTopicConsumer(channel, conf)
	default:
		return NewSimpleConsumer(channel, conf)
	}
}
