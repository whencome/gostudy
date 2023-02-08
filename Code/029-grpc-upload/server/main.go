package main

import (
	"demo029/pb"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	grpc "google.golang.org/grpc"
)

var _ pb.UploadServer = (*UploadServerImpl)(nil)

type UploadServerImpl struct {
	pb.UnimplementedUploadServer
}

// Upload 接收客户端上传文件请求
func (s *UploadServerImpl) Upload(stream pb.Upload_UploadServer) error {
	var file *os.File
	fileName := ""
	fileSize := int64(0)
	wroteSize := int64(0)
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// 正常接收完毕
				break
			}
		}
		// 初次接收请求，保存文件名以及文件大小
		if fileName == "" {
			fileName = req.Name
			fileSize = req.FileSize
			// 在当前目录创建文件
			file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModeAppend)
			if err != nil {
				return fmt.Errorf("create file: %w", err)
			}
			defer file.Close()
		}
		// 不接受空数据
		if len(req.Data) == 0 {
			return fmt.Errorf("empty data")
		}
		// 追加写入文件内容
		n, err := file.Write(req.Data)
		if err != nil {
			return fmt.Errorf("write file: %w", err)
		}
		wroteSize += int64(n)
		if wroteSize > fileSize {
			return fmt.Errorf("data size exceeded: expected %d bytes, got %d bytes", fileSize, wroteSize)
		}
	}
	resp := &pb.UploadResponse{
		Path: fileName,
	}
	return stream.SendAndClose(resp)
}

func main() {
	listen, err := net.Listen("tcp4", ":8009")
	if err != nil {
		log.Panicf("listen 8009 failed: %v", err)
	}

	svr := grpc.NewServer()
	pb.RegisterUploadServer(svr, &UploadServerImpl{})
	if err = svr.Serve(listen); err != nil {
		log.Panicf("Serve failed: %v", err)
	}
}
