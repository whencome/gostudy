package def

// 定义消息消费函数
type ConsumeFunc func([]byte) error

// 定义模式
const (
	// Simple模式
	ModeSimple string = "simple"
	// Worker模式
	ModeWork string = "work"
	// Publish模式
	ModePublish string = "publish"
	// Routing模式
	ModeRouting string = "routing"
	// Topic模式
	ModeTopic string = "topic"
)

// Config 定义rabbitmq配置
type Config struct {
	// MQ连接地址
	MQUrl string
	// 模式
	Mode string
	// 队列名称
	QueueName string
	// 交换机名称
	Exchange string
	// bind Key 名称
	Key string
	// 是否持久化
	Durable bool
	// 是否自动确认消息
	AutoAck bool
}
