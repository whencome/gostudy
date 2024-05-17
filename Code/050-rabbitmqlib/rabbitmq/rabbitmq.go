package rabbitmq

import (
	"errors"

	"github.com/streadway/amqp"

	"demo050/rabbitmq/consumer"
	"demo050/rabbitmq/def"
	"demo050/rabbitmq/producer"
)

// rabbitMQ结构体
type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	conf     *def.Config
	consumer consumer.Consumer
	producer producer.Producer
}

// 创建结构体实例
func New(conf *def.Config) *RabbitMQ {
	return &RabbitMQ{
		conf: conf,
	}
}

// Close 断开channel 和 connection
func (r *RabbitMQ) Close() {
	if r.consumer != nil {
		r.consumer.Stop()
	}
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}

// Open 打开连接，并获取channel
func (r *RabbitMQ) Open() error {
	var err error
	// 获取connection
	r.conn, err = amqp.Dial(r.conf.MQUrl)
	if err != nil {
		return err
	}
	// 获取channel
	r.channel, err = r.conn.Channel()
	if err != nil {
		return err
	}
	// 获取producer & consumer
	r.producer = producer.New(r.channel, r.conf)
	r.consumer = consumer.New(r.channel, r.conf)
	// 返回结果
	return nil
}

// Publish 发布消息
func (r *RabbitMQ) Publish(message []byte) error {
	if r.producer == nil {
		return errors.New("producer not initialized")
	}
	return r.producer.Publish(message)
}

// Consume 消费消息
func (r *RabbitMQ) Consume(f def.ConsumeFunc) error {
	if r.producer == nil {
		return errors.New("consumer not initialized")
	}
	return r.consumer.Consume(f)
}
