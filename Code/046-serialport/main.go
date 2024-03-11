package main

import (
	"log"

	"github.com/jacobsa/go-serial/serial"
)

// tutorial: https://developer.aliyun.com/article/1378633
func main() {
	// 配置串口参数
	options := serial.OpenOptions{
		PortName:        "COM1",
		BaudRate:        115200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	// 打开串口
	port, err := serial.Open(options)
	if err != nil {
		log.Printf("open port fail: %v", err)
		return
	}
	log.Println("port opened")
	// 关闭串口
	defer port.Close()

	// 发送数据
	sendData := []byte("Hello, Serial!")
	n, err := port.Write(sendData)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sent %d bytes: %v", n, sendData)

	// 接收数据
	buf := make([]byte, 128)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Received %d bytes: %v", n, buf[:n])

}
