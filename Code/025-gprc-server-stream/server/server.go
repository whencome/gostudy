package main

import (
	pb "demo025/pb/ecommerce"
	"fmt"
	"net"

	grpc "google.golang.org/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// compile time type check
var _ pb.OrderManagementServer = &OrderManagementImpl{}

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

// ------------------ business implementation ------------------------ //

// OrderManagementImpl implement pb.OrderManagementServer
type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

// SearchOrders business implementation
func (s *OrderManagementImpl) SearchOrders(query *wrapperspb.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for _, order := range orders {
		err := stream.Send(&order)
		if err != nil {
			return fmt.Errorf("error send order: %v", err)
		}
	}
	return nil
}

// ------------------ service entry ------------------- //

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
