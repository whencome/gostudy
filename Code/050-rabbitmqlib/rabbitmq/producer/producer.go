package producer

import (
	"demo050/rabbitmq/def"

	"github.com/streadway/amqp"
)

// Producer 定义一个生产者接口
type Producer interface {
	Publish(message []byte) error
}

// New 创建一个新的producer
func New(channel *amqp.Channel, conf *def.Config) Producer {
	switch conf.Mode {
	case def.ModePublish:
		return NewPublishProducer(channel, conf)
	case def.ModeRouting:
		return NewRoutingProducer(channel, conf)
	case def.ModeTopic:
		return NewTopicProducer(channel, conf)
	default:
		return NewSimpleProducer(channel, conf)
	}
}
