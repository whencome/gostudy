package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	addr := "0.0.0.0:3031"
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		fmt.Printf("listen to %v failed: %s\n", addr, err)
		return
	}
	fmt.Printf("tcp server started on %s\n", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen failed: ", err)
			return
		}
		// 处理链接
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read data failed: ", err)
		return
	}
	fmt.Println("GOT: ", string(line))
	conn.Write(line)
}
