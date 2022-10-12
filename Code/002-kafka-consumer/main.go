package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	// 创建消费者
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("create kafka consumer failed: ", err)
		return
	}

	// 获取topic分区列表
	partitionList, err := consumer.Partitions("test_topic")
	if err != nil {
		fmt.Println("get partitions failed: ", err)
		return
	}
	fmt.Println("partitions: ", partitionList)

	// 异步从每个分区消费消息
	wg := new(sync.WaitGroup)
	for _, partition := range partitionList {
		partitionConsumer, err := consumer.ConsumePartition("test_topic", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("consume partition failed: ", err)
			continue
		}
		defer partitionConsumer.AsyncClose()

		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("timestamp => %s, partition => %d, offset => %d, topic => %s, key => %s, value => %s\n", msg.Timestamp.String(), msg.Partition, msg.Offset, msg.Topic, msg.Key, msg.Value)
			}
			wg.Done()
		}(partitionConsumer)
	}
	wg.Wait()
}
