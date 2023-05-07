package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {
	// 创建一个缓存对象，缓存时间5分钟，每10秒清理一次过期缓存
	c := cache.New(time.Minute*5, time.Second*10)
	// 设置缓存，缓存5秒钟
	c.Set("foo", 12, time.Second*5)
	v, found := c.Get("foo")
	if !found {
		fmt.Printf("foo not found\n")
	} else {
		fmt.Printf("foo = %v\n", v)
	}
	time.Sleep(time.Second * 12)
	fmt.Println("have waited for 12 seconds...")
	v, found = c.Get("foo")
	if !found {
		fmt.Printf("foo not found\n")
	} else {
		fmt.Printf("foo = %v\n", v)
	}
}
