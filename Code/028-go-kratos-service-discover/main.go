package main

import (
	"context"
	"log"
	"time"

	etcdregistry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	etcdclient "go.etcd.io/etcd/client/v3"

	pb "demo028/api/helloworld/v1"
)

func main() {
	// 获取etcd连接
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{
			"127.0.0.1:20000",
			"127.0.0.1:20002",
			"127.0.0.1:20004",
		},
	})
	if err != nil {
		log.Panicf("get etcd client failed: %v", err)
	}
	// 使用etcd作为服务注册与发现服务中心
	r := etcdregistry.New(client)

	// 获取grpc连接
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	connGrpc, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("discovery:///helloworld"), // 服务发现
		grpc.WithDiscovery(r),                        // 服务中心
	)
	if err != nil {
		log.Panicf("get grpc connection failed: %v", err)
	}
	defer connGrpc.Close()

	// 服务调用
	userClient := pb.NewUserClient(connGrpc)
	userResp, err := userClient.QueryUser(ctx, &pb.QueryUserRequest{Id: 1})
	if err != nil {
		log.Panicf("query user info failed: %v", err)
	}
	// print response
	log.Printf("User Info:\n-----------\nID: %d\nName: %s\nAge: %d\n", userResp.Id, userResp.Name, userResp.Age)
}
