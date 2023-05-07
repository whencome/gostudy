package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {

	c := cache.New(time.Minute*5, time.Second*10)
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
