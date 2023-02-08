// for more instrcution, see: https://www.lixueduan.com/posts/grpc/05-interceptor/
package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// ------------------ unary interceptor ---------------------//

func ClientUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// pre process
	start := time.Now()
	// processing
	err := invoker(ctx, method, req, reply, cc, opts...)
	// post process
	end := time.Now()
	// log
	log.Printf("CLIENT Unary Rpc: %s; req: %+v;  resp: %+v; start: %s; end: %s; err: %v", method, req, reply, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	// return
	return err
}

// ------------------ stream interceptor ---------------------//

type WrappedClientStream struct {
	grpc.ClientStream
}

func NewWrappedClientStream(c grpc.ClientStream) *WrappedClientStream {
	return &WrappedClientStream{c}
}

func (s *WrappedClientStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return s.ClientStream.RecvMsg(m)
}

func (s *WrappedClientStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return s.ClientStream.SendMsg(m)
}

func ClientStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return NewWrappedClientStream(s), nil
}
