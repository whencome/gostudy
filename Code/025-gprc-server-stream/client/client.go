package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "demo025/pb/ecommerce"

	grpc "google.golang.org/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cli := pb.NewOrderManagementClient(conn)

	stream, err := cli.SearchOrders(ctx, &wrapperspb.StringValue{Value: "1001"})
	if err != nil {
		panic(err)
	}

	for {
		order, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("recv from server finished")
				break
			}
			fmt.Printf("recv err: %v\n", err)
			break
		}
		// print order info
		fmt.Printf("recv: %#v\n", order)
	}

	fmt.Println("\n----------- over -------------\n")
}
