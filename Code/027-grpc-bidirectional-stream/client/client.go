package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "demo027/pb/ecommerce"

	grpc "google.golang.org/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	orderIds := []string{"1001", "1002", "1003", "1004"}

	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	cli := pb.NewOrderManagementClient(conn)
	stream, err := cli.ProcessOrders(ctx)
	if err != nil {
		panic(err)
	}

	// send request
	go func() {
		for _, orderId := range orderIds {
			fmt.Println("--- send orderId: ", orderId)
			err = stream.Send(&wrapperspb.StringValue{Value: orderId})
			if err != nil {
				fmt.Printf("send error: %v", err)
				continue
			}
		}
	}()

	// handle response
	for {
		quit := false
		select {
		case <-ctx.Done():
			quit = true
			break
		default:
			fmt.Println("### recv from server ####")
			shp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Printf("recv error: %v\n", err)
				quit = true
				break
			}
			fmt.Printf("recvd : %v\n\n", *shp)
		}
		if quit {
			break
		}
	}

	fmt.Println("\n ------------ over ------------ \n")
}
