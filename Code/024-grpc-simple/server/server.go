package main

import (
	"context"
	pb "demo024/pb/ecommerce"
	"net"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// 编译期类型检查
var _ pb.OrderManagementServer = &OrderManagementImpl{}

// ------------- business implement ------------- //

var orders = make(map[string]pb.Order)

type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

// GetOrder 实现服务接口
func (s *OrderManagementImpl) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	order, exists := orders[orderId.Value]
	if exists {
		return &order, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "order not found : ", orderId.Value)
}

// -------------- service entry ----------------//
func main() {
	// create server
	svr := grpc.NewServer()

	// resister service to server
	pb.RegisterOrderManagementServer(svr, &OrderManagementImpl{})

	// listen and start server
	lis, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}
	if err = svr.Serve(lis); err != nil {
		panic(err)
	}

}
