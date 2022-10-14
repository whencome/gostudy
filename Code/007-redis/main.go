package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis failed: ", err)
		return
	}
	defer conn.Close()

	// 授权
	_, err = conn.Do("AUTH", "123456")
	if err != nil {
		fmt.Println("auth failed: ", err)
		return
	}

	// 设置值
	_, err = conn.Do("SET", "name", "whencome")
	if err != nil {
		fmt.Println("set value failed: ", err)
		return
	}

	// 取值
	v, err := conn.Do("GET", "name")
	if err != nil {
		fmt.Println("get value failed: ", err)
		return
	}
	fmt.Printf("got value: %s\n", string(v.([]byte)))
}
