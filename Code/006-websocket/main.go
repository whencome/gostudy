/**
 * websocket教程参考：https://zhuanlan.zhihu.com/p/455635795
 */
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// 定义升级程序
var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func wsEntry(w http.ResponseWriter, r *http.Request) {
	// 检查请求来源，是否允许跨域请求
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// 升级连接
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("upgrade fail: %s\n", err)
		return
	}

	// 处理请求
	handleWsConn(ws)
}

func handleWsConn(conn *websocket.Conn) error {
	// 发送欢饮消息
	_ = conn.WriteMessage(1, []byte("welcome to websocket demo"))
	// 这里是无限循环，用于和用户交互
	for {
		// 读取客户端消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read message failed: ", err)
			continue
		}
		// 打印消息
		fmt.Printf("msg_type = %v, msg = %v\n", msgType, string(msg))

		// 向客户端发送消息
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println("send message failed: ", err)
		}
		//
		_ = conn.WriteMessage(2, []byte(time.Now().String()))
	}
}

func main() {
	http.HandleFunc("/ws", wsEntry)
	if err := http.ListenAndServe(":9099", nil); err != nil {
		fmt.Printf("run http failed: %s\n", err)
	}
}
