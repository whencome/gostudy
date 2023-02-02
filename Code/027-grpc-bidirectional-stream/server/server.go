package main

import (
	pb "demo027/pb/ecommerce"
	"fmt"
	"io"
	"net"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// define a order list
var orders = map[string]*pb.Order{
	"1001": &pb.Order{
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
	"1002": &pb.Order{
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
	"1003": &pb.Order{
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

type OrderManagementImpl struct {
	pb.UnsafeOrderManagementServer
}

func (s *OrderManagementImpl) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	for {
		fmt.Println("---- read from client")
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("recv error: %v", err)
			return err
		}

		fmt.Printf("request process order %s \n", req.Value)

		order, exists := orders[req.Value]
		if !exists {
			fmt.Printf("order %s not exists\n", req.Value)
			return status.Errorf(codes.NotFound, "order not found : ", req.Value)
		}

		shp := &pb.Shipment{
			Id:     req.Value,
			Status: "processed",
			Order:  order,
		}
		err = stream.Send(shp)
		if err != nil {
			fmt.Printf("send error: %v", err)
			return err
		}
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
