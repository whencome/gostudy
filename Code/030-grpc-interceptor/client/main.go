package main

import (
	"context"
	"demo030/interceptor"
	"demo030/pb"
	"io"
	"log"
	"sync"

	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor.ClientUnaryInterceptor), grpc.WithStreamInterceptor(interceptor.ClientStreamInterceptor))
	if err != nil {
		log.Panicf("grpc dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	unarySay(client)
	streamSay(client)

	log.Println("---------- over -----------")
}

func unarySay(client pb.EchoClient) {
	uReq := &pb.UnarySayRequest{
		Req: "unary-req",
	}
	uResp, err := client.UnarySay(context.Background(), uReq)
	if err != nil {
		log.Printf("unary say err: %v\n", err)
		return
	}
	log.Println("unary say resp: ", uResp.Resp)
}

func streamSay(client pb.EchoClient) {
	sReqs := []*pb.StreamSayRequest{
		{Req: "stream-req-1"},
		{Req: "stream-req-2"},
		{Req: "stream-req-3"},
	}
	stream, err := client.StreamSay(context.Background())
	if err != nil {
		log.Printf("stream say : get stream err: %v\n", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, r := range sReqs {
			log.Println("--> send :", r.Req)
			err := stream.Send(r)
			if err != nil {
				log.Printf("stream say : send err: %v\n", err)
				break
			}
		}
		stream.CloseSend()
	}()

	go func() {
		defer wg.Done()
		for {
			r, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Printf("stream say : recv err: %v\n", err)
				break
			}
			log.Println("<-- recv :", r.Resp)
		}
	}()
	wg.Wait()
}
