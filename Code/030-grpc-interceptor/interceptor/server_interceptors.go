// for more instrcution, see: https://www.lixueduan.com/posts/grpc/05-interceptor/
package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// ------------------ unary interceptor ---------------------//

func ServerUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// pre process
	start := time.Now()
	// processing
	reply, err := handler(ctx, req)
	// post process
	end := time.Now()
	// log
	log.Printf("SERVER Unary Rpc: %s; req: %+v;  resp: %+v; start: %s; end: %s; err: %v", info.FullMethod, req, reply, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	// return
	return reply, err
}

// ------------------ stream interceptor ---------------------//

type WrappedServerStream struct {
	grpc.ServerStream
}

func NewWrappedServerStream(s grpc.ServerStream) *WrappedServerStream {
	return &WrappedServerStream{s}
}

func (s *WrappedServerStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return s.ServerStream.RecvMsg(m)
}

func (s *WrappedServerStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return s.ServerStream.SendMsg(m)
}

func ServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, NewWrappedServerStream(ss))
}
