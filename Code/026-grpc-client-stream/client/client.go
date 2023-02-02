package main

import (
	"context"
	"fmt"
	"time"

	pb "demo026/pb/ecommerce"

	grpc "google.golang.org/grpc"
)

// define a order list
var orders = []pb.Order{
	{
		Id: "1001",
		Items: []string{
			"1001 - item 1",
			"1001 - item 2",
			"1001 - item 3",
		},
		Description: "description of order 1001",
		Price:       99.15,
		Destination: "destination of order 1001",
	},
	{
		Id: "1002",
		Items: []string{
			"1002 - item 1",
			"1002 - item 2",
			"1002 - item 3",
		},
		Description: "description of order 1002",
		Price:       287.15,
		Destination: "destination of order 1002",
	},
	{
		Id: "1003",
		Items: []string{
			"1003 - item 1",
			"1003 - item 2",
			"1003 - item 3",
		},
		Description: "description of order 1003",
		Price:       1899,
		Destination: "destination of order 1003",
	},
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cli := pb.NewOrderManagementClient(conn)
	stream, err := cli.UpdateOrders(ctx)
	for _, order := range orders {
		err = stream.Send(&order)
		if err != nil {
			panic(err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println("response data: ", resp.Value)
}
