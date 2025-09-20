package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
    clientConfig := constant.ClientConfig{
        TimeoutMs:      5000,
        ListenInterval: 10000,
        Endpoint:       "localhost:8848",
        NamespaceId:    "public",
    }

    serverConfigs := []constant.ServerConfig{
        {
            IpAddr: "127.0.0.1",
            Port:   8848,
        },
    }

    namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
    if err != nil {
        log.Panic(err)
    }

    // 查询服务实例列表
    result, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
        ServiceName: "demo-service",
        GroupName:   "DEFAULT_GROUP",
    })

    if err != nil {
        log.Fatal("查询服务实例失败:", err)
    }

    fmt.Println("发现的服务实例:")
    for _, instance := range result {
        fmt.Printf("IP: %s, Port: %d, Weight: %.2f, Metadata: %v, Healthy: %t\n",
            instance.Ip, instance.Port, instance.Weight, instance.Metadata, instance.Healthy)
    }

    // 监听服务变化（可选）
    namingClient.Subscribe(&vo.SubscribeParam{
        ServiceName: "demo-service",
        GroupName:   "DEFAULT_GROUP",
        SubscribeCallback: func (services []model.Instance, err error)  {
            if err != nil {
                log.Println("监听服务变化失败:", err)
            } else {
                fmt.Println("服务变化:")
                for _, service := range services {
                    fmt.Printf("IP: %s, Port: %d, Weight: %.2f, Metadata: %v, Healthy: %t\n",
                        service.Ip, service.Port, service.Weight, service.Metadata, service.Healthy)
                }
            }
        },
    })

    // 保持监听
    time.Sleep(60 * time.Second)
}