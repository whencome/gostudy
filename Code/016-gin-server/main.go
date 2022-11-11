package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func initRoute(r *gin.Engine) {
	r.GET("/greet", func(c *gin.Context) {
		name := c.DefaultQuery("name", "James")
		fmt.Fprintf(c.Writer, "hello %s", name)
	})
}

func HandleSignal(sigChan chan os.Signal, releaseFunc func()) {
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP:
			// 释放相关资源
			if releaseFunc != nil {
				releaseFunc()
			}
			// 等待1秒再退出
			time.Sleep(1 * time.Second)
			os.Exit(0)
			return
		// 非退出信号不作处理，继续等待
		default:
			continue
		}
	}
}

var sigChan chan os.Signal = make(chan os.Signal)
var svr *HTTPServer

func main() {
	opts := &Options{
		Port: 8910,
		Mode: ModeDebug,
	}
	svr = New(opts)
	svr.Init(initRoute)
	svr.PreInit(func() error {
		log.Println("--------- pre init ---------")
		return nil
	})
	svr.PostInit(func() error {
		log.Println("--------- post init ---------")
		return nil
	})
	svr.PreStop(func() error {
		log.Println("--------- pre stop ---------")
		return nil
	})
	svr.PostStop(func() error {
		log.Println("--------- post stop ---------")
		return nil
	})
	svr.Start()

	HandleSignal(sigChan, func() {
		svr.Stop()
	})
}
