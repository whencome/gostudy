/**
 * 已关闭的通道不能再写值，但是是可以继续读取的，当没有值的时候，会无阻塞获取零值，一般用于并发控制（通道不用于值的实际传递）。
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	exitChan := make(chan struct{})

	go func() {
		<-exitChan
		fmt.Println("-- exit a --")
	}()
	go func() {
		<-exitChan
		fmt.Println("-- exit b --")
	}()
	go func() {
		<-exitChan
		fmt.Println("-- exit c --")
	}()
	go func() {
		<-exitChan
		fmt.Println("-- exit d --")
	}()

	fmt.Println("----- game begin -----")
	time.Sleep(time.Second * 5)
	close(exitChan)
	fmt.Println("----- game playing -----")
	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		<-exitChan
		fmt.Println("$$ read from closed chan $$")
	}
	fmt.Println("----- game over-----")
}
