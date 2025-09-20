package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func waitForNacosReady(namingClient naming_client.INamingClient) error {
	log.Println("等待 Nacos 服务端连接...")
	for i := 0; i < 30; i++ { // 最多等待 30 秒
		_, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
			ServiceName: "nacos.discovery.test.service", // 一个测试服务名，无需真实存在
		})
		if err == nil {
			log.Println("Nacos 客户端已连接并就绪")
			return nil
		}
		log.Printf("连接中... 错误: %v", err)
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("等待 Nacos 连接超时")
}

func main() {
	// 1. 创建客户端配置
	clientConfig := constant.ClientConfig{
		TimeoutMs:      5000,
		ListenInterval: 10000,
		Endpoint:       "localhost:8848", // Nacos 服务端地址
		NamespaceId:    "public",         // 命名空间 ID，public 可以留空或填 public
		// LogLevel:       "debug",          // 启用 debug 日志
		// LogDir:         "./log",
		// CacheDir:       "./cache",
	}

	// 2. 创建服务端配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	// 3. 创建服务发现客户端
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Panic("创建 Nacos 客户端失败:", err)
	}

	// ✅ 等待客户端连接到 Nacos Server
	log.Println("正在连接 Nacos Server...")

	// 创建 namingClient 后
	err = waitForNacosReady(namingClient)
	if err != nil {
		log.Fatal("Nacos 连接失败:", err)
	}
	log.Println("✅ 已连接到 Nacos Server...")

	// ✅ 现在可以安全注册服务
	ip, _ := getLocalIP()
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        8080,
		ServiceName: "demo-service",
		GroupName:   "DEFAULT_GROUP",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"version": "v1.0.0"},
	})

	if !success || err != nil {
		log.Fatalf("服务注册失败: %v, error: %v", success, err)
	}

	log.Println("服务注册成功:", ip, ":8080")

	// 保持运行，模拟服务在线
	time.Sleep(120 * time.Second)

	// 6. （可选）服务注销
	_, _ = namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        8080,
		ServiceName: "demo-service",
		GroupName:   "DEFAULT_GROUP",
	})

	fmt.Println("服务已从 Nacos 注销")
}

// 获取本机 IP 地址
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "127.0.0.1", nil
}
