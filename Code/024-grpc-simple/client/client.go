package main

import (
	"context"
	pb "demo024/pb/ecommerce"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	// dial to server
	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// get a client
	cli := pb.NewOrderManagementClient(conn)

	// create a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call rpc service to get order
	order, err := cli.GetOrder(ctx, &wrapperspb.StringValue{Value: "100"})
	if err != nil {
		panic(err)
	}

	log.Printf("get order: %#v\n", order)
}
