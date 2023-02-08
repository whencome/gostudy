package main

import (
	"context"
	"demo030/interceptor"
	"demo030/pb"
	"fmt"
	"io"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type EchoServerImpl struct {
	pb.UnimplementedEchoServer
}

func (s *EchoServerImpl) UnarySay(ctx context.Context, req *pb.UnarySayRequest) (*pb.UnarySayResponse, error) {
	return &pb.UnarySayResponse{
		Resp: fmt.Sprintf("SVR RESP: %s", req.Req),
	}, nil
}

func (s *EchoServerImpl) StreamSay(stream pb.Echo_StreamSayServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("recv err : %v", err)
			break
		}
		resp := &pb.StreamSayResponse{
			Resp: fmt.Sprintf("SVR RESP: %s", req.Req),
		}
		err = stream.Send(resp)
		if err != nil {
			log.Printf("send err : %v", err)
			break
		}
	}
	return nil
}

func main() {
	svr := grpc.NewServer(grpc.UnaryInterceptor(interceptor.ServerUnaryInterceptor), grpc.StreamInterceptor(interceptor.ServerStreamInterceptor))
	pb.RegisterEchoServer(svr, &EchoServerImpl{})

	lis, err := net.Listen("tcp4", ":8009")
	if err != nil {
		log.Panicf("listen 8009 failed: %v", err)
	}
	if err = svr.Serve(lis); err != nil {
		log.Panicf("serve failed: %v", err)
	}
}
