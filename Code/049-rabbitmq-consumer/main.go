package main

func main() {
	conf := &Config{
		MQUrl:     "amqp://root:root@127.0.0.1:5672/",
		QueueName: "simple_queue",
		Exchange:  "",
		Key:       "",
		Durable:   true,
	}
	rabbitmq := NewRabbitMQSimple(conf)
	rabbitmq.ConsumeSimple()
}
