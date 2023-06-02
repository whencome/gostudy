package main

import (
    "bufio"
    "bytes"
    "flag"
    "fmt"
    gnet "github.com/panjf2000/gnet/v2"
    "log"
    "strconv"
)

type echoServer struct {
    gnet.BuiltinEventEngine

    eng       gnet.Engine
    addr      string
    multicore bool
}

func (es *echoServer) OnBoot(eng gnet.Engine) gnet.Action {
    es.eng = eng
    log.Printf("echo server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
    return gnet.None
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
    reader := bufio.NewReader(c)
    // 读取命令长度
    line1, _, err := reader.ReadLine()
    if err != nil {
        log.Panicf("read line 1 err: %v", err)
        return gnet.Close
    }
    log.Printf("line1 => %s", string(line1))
    // 读取命令
    line2, _, err := reader.ReadLine()
    if err != nil {
        log.Panicf("read line 2 err: %v", err)
        return gnet.Close
    }
    log.Printf("line2 => %s", string(line2))
    // 读取正文内容
    // 正文长度
    line3, _, err := reader.ReadLine()
    if err != nil {
        log.Panicf("read line 3 err: %v", err)
        return gnet.Close
    }
    log.Printf("line3 => %s", string(line3))
    // 正文内容
    line4, _, err := reader.ReadLine()
    if err != nil {
        log.Panicf("read line 4 err: %v", err)
        return gnet.Close
    }
    log.Printf("line4 => %s", string(line4))

    //buf, _ := c.Next(100)
    resp := bytes.Buffer{}
    resp.Write([]byte(strconv.Itoa(len(line4))))
    resp.Write([]byte{'\n'})
    resp.Write(line4)
    resp.Write([]byte{'\n'})
    c.Write(resp.Bytes())
    return gnet.None
}

func main() {
    var port int
    var multicore bool

    // Example command: go run echo.go --port 9000 --multicore=true
    flag.IntVar(&port, "port", 9000, "--port 9000")
    flag.BoolVar(&multicore, "multicore", false, "--multicore true")
    flag.Parse()
    echo := &echoServer{addr: fmt.Sprintf("tcp://:%d", port), multicore: multicore}
    log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore)))
}
