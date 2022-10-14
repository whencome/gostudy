package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis failed: ", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("AUTH", "123456")
	if err != nil {
		fmt.Println("auth failed: ", err)
		return
	}

	// 1. expire after N seconds
	_, err = conn.Do("SET", "key1", "value1")
	if err != nil {
		fmt.Println("set key1 failed: ", err)
		return
	}
	v1, err := redis.String(conn.Do("GET", "key1"))
	if err != nil {
		fmt.Println("get key1 failed: ", err)
		return
	}
	fmt.Println("key1 = ", v1)
	// set expire after 3 seconds
	// ignore errors from here
	_, _ = conn.Do("EXPIRE", "key1", 3)
	time.Sleep(4 * time.Second)
	v1, _ = redis.String(conn.Do("GET", "key1"))
	fmt.Println("after 4 seconds, key1 = ", v1)

	// 2. expire at fixed time
	_, _ = conn.Do("SET", "key2", "value2")
	v2, _ := redis.String(conn.Do("GET", "key2"))
	fmt.Println("key2 = ", v2)
	expireTime := time.Now().Add(time.Second * 3)
	_, _ = conn.Do("EXPIREAT", "key2", expireTime.Unix())
	time.Sleep(4 * time.Second)
	v1, _ = redis.String(conn.Do("GET", "key2"))
	fmt.Printf("after %s, key2 = %s\n", expireTime.String(), v1)
}
