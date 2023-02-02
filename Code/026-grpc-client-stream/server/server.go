package main

import (
	pb "demo026/pb/ecommerce"
	"fmt"
	"io"
	"net"
	"strings"

	grpc "google.golang.org/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.OrderManagementServer = &OrderManagementImpl{}

type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

func (s *OrderManagementImpl) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	buf := strings.Builder{}
	for {
		order, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&wrapperspb.StringValue{Value: buf.String()})
			}
			return fmt.Errorf("recv error: %w", err)
		}
		fmt.Printf("recvd: %#v\n", *order)
		buf.WriteString(order.Id)
		buf.WriteString(",")
	}
	return nil
}

func main() {
	svr := grpc.NewServer()
	pb.RegisterOrderManagementServer(svr, &OrderManagementImpl{})
	lis, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}
	if err = svr.Serve(lis); err != nil {
		panic(err)
	}
}
