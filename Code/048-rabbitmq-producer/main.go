package main

import (
	"fmt"
	"time"
)

func main() {
	rabbitmq := NewRabbitMQSimple("simple_queue")
	rabbitmq.PublishSimple(fmt.Sprintf("a message from go at %s", time.Now().Format("2006-01-02 15:04:05")))
	fmt.Println("--- OK ---")
}
