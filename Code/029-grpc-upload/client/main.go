package main

import (
	"context"
	"demo029/pb"
	"io"
	"log"
	"os"

	grpc "google.golang.org/grpc"
)

func main() {
	filePath := "light-space-image.jpg"

	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil {
		log.Panicf("dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewUploadClient(conn)
	resp, err := uploadFile(client, filePath)
	if err != nil {
		log.Panicf("upload fail: %v", err)
	}
	log.Println("upload to: ", resp.Path)
}

func uploadFile(client pb.UploadClient, p string) (*pb.UploadResponse, error) {
	fileInfo, err := os.Stat(p)
	if err != nil {
		log.Printf("stat err: %v\n", err)
		return nil, err
	}
	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()
	buf := make([]byte, 4096)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.Upload(ctx)
	if err != nil {
		log.Printf("get stream err: %v\n", err)
		return nil, err
	}

	f, err := os.OpenFile(p, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Printf("open file err: %v\n", err)
		return nil, err
	}
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		req := &pb.UploadRequest{
			Name:     fileName,
			FileSize: fileSize,
			Data:     buf[0:n],
		}
		err = stream.Send(req)
		if err != nil {
			log.Printf("send err: %v\n", err)
			return nil, err
		}
	}
	return stream.CloseAndRecv()
}
