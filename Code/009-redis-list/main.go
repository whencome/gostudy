package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// 获取redis连接
func getRedisConn() (*redis.Client, error) {
	options := &redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	}
	client := redis.NewClient(options)
	// check connection
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getListData(cli *redis.Client, key string) ([]string, error) {
	len, err := cli.LLen(key).Result()
	if err != nil {
		return nil, err
	}
	if len == 0 {
		return []string{}, nil
	}
	return cli.LRange(key, 0, len).Result()
}

func printBooks(cli *redis.Client) {
	data, err := getListData(cli, "books")
	if err != nil {
		fmt.Println("get books from redis failed: ", err)
		return
	}
	fmt.Printf("books: %+v\n", data)
}

func main() {
	redisCli, err := getRedisConn()
	if err != nil {
		fmt.Println("connect redis failed: ", err)
		return
	}

	// remove key books
	redisCli.Del("books")

	err = redisCli.LPush("books", "C语言从入门到精通", "疯狂java讲义", "图解HTTP").Err()
	if err != nil {
		fmt.Println("set books failed: ", err)
		return
	}
	printBooks(redisCli)

	redisCli.RPush("books", "编程的逻辑")
	printBooks(redisCli)

	redisCli.RPush("books", "Go微服务实战", "具体数学")
	printBooks(redisCli)

	book := redisCli.LPop("books").Val()
	fmt.Println("LPop a book: ", book)
	printBooks(redisCli)

	book = redisCli.RPop("books").Val()
	fmt.Println("RPop a book: ", book)
	printBooks(redisCli)

}
