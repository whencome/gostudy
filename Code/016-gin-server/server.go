package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ModeDebug   = "debug"
	ModeRelease = "release"
)

type Options struct {
	Port int    `json:"port" yaml:"port" toml:"port"` // 服务端口
	Mode string `json:"mode" yaml:"mode" toml:"mode"` // 运行模式, debug or release
}

type InitFunc func(r *gin.Engine)
type HookFunc func() error

// DefaultOptions create a default options
func DefaultOptions() *Options {
	return &Options{
		Port: 8080,
		Mode: ModeRelease,
	}
}

// 定义一个HTTP服务
type HTTPServer struct {
	// 服务是否正在允许
	running bool
	// gin engine
	engine *gin.Engine
	// http服务信息
	svr *http.Server
	// 站点配置
	options *Options
	// 初始化方法
	initFunc InitFunc
	// 初始化前后执行的方法
	preInitFunc  HookFunc
	postInitFunc HookFunc
	// 服务通知之后执行的方法
	preStopFunc  HookFunc
	postStopFunc HookFunc
}

// 创建一个server
func New(options *Options) *HTTPServer {
	if options == nil {
		options = DefaultOptions()
	}
	s := &HTTPServer{
		running: false,
		engine:  gin.Default(),
		svr:     nil,
		options: options,
	}
	// 初始化http server
	s.svr = &http.Server{
		Addr:    fmt.Sprintf(":%d", options.Port),
		Handler: s.engine,
	}
	return s
}

// Init 初始化路由
func (s *HTTPServer) Init(f InitFunc) {
	s.initFunc = f
}

func (s *HTTPServer) PreInit(f HookFunc) {
	s.preInitFunc = f
}

func (s *HTTPServer) PostInit(f HookFunc) {
	s.postInitFunc = f
}

func (s *HTTPServer) PreStop(f HookFunc) {
	s.preStopFunc = f
}

func (s *HTTPServer) PostStop(f HookFunc) {
	s.postStopFunc = f
}

func (s *HTTPServer) execHook(f HookFunc, action string) {
	if f == nil {
		return
	}
	e := f()
	if e != nil {
		log.Printf("%s fail: %s\n", action, e)
	}
}

func (s *HTTPServer) execHookMustSucc(f HookFunc, action string) {
	if f == nil {
		return
	}
	e := f()
	if e != nil {
		log.Panicf("%s fail %s\n", action, e)
	}
}

// 判断服务是否可以运行
func (s *HTTPServer) IsRunnable() bool {
	return !s.running
}

// 启动服务
func (s *HTTPServer) Start() {
	if !s.IsRunnable() {
		log.Println("can not start server, it's probably already started")
		return
	}
	// 设置运行模式
	if s.options.Mode != ModeDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	// 执行初始化之前的操作
	s.execHookMustSucc(s.postInitFunc, "pre init")
	if s.initFunc != nil {
		s.initFunc(s.engine)
	}
	s.execHookMustSucc(s.postInitFunc, "post init")
	// 启动http服务
	s.running = true
	go func() {
		if err := s.svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("run server failed: %s \n", err)
		}
	}()
	log.Printf("http service started on %s", s.svr.Addr)
}

// 停止服务
func (s *HTTPServer) Stop() {
	s.execHook(s.preStopFunc, "prepare stop server")
	log.Println("start to shutdown http server")
	if err := s.svr.Shutdown(context.Background()); err != nil {
		log.Printf("shutdown server failed: %s \n", err)
		return
	}
	s.running = false
	s.execHook(s.postStopFunc, "")
	log.Println("http server closed")
}
