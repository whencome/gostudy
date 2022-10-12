package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	addr := "127.0.0.1:3031"
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		fmt.Printf("connect to %v failed: %s", addr, err)
		return
	}
	defer conn.Close()
	// send data
	msg := "hello from " + time.Now().String()
	n, err := conn.Write([]byte(msg + "\n")) // 必须要加换行符才会直接发送到服务端
	if err != nil {
		fmt.Println("send message failed: ", err)
		return
	}
	fmt.Printf("send message: %s; %d bytes send\n", msg, n)
	// read data
	servResp, err := readData(conn)
	if err != nil {
		fmt.Println("read from server failed: ", err)
		return
	}
	fmt.Println("server response: ", string(servResp))
}

func readData(conn net.Conn) ([]byte, error) {
	result := make([]byte, 0)
	for {
		var data [1024]byte
		n, err := conn.Read(data[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return result, err
		}
		result = append(result, data[:n]...)
	}
	return result, nil
}
