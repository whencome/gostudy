package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	endPoints := []string{
		"http://127.0.0.1:20000",
		"http://127.0.0.1:20002",
		"http://127.0.0.1:20004",
	}
	etcdx := NewEtcdx(endPoints, 3)
	defer etcdx.Close()

	fmt.Println("demo running...")

	go etcdx.Watch("name", printChange)
	go etcdx.Watch("age", printChange)

	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func(wg *sync.WaitGroup, x *Etcdx) {
		for i := 0; i < 10; i++ {
			x.Put("name", fmt.Sprintf("test-%d", i))
			x.Put("age", fmt.Sprintf("%d", 20+i))
			time.Sleep(time.Second * 1)
		}
		wg.Done()
	}(wg, etcdx)
	go func(wg *sync.WaitGroup, x *Etcdx) {
		for i := 0; i < 15; i++ {
			name, _ := x.Get("name")
			age, _ := x.Get("age")
			fmt.Printf("name = %s, age = %s\n", name, age)
			time.Sleep(time.Second * 1)
		}
		wg.Done()
	}(wg, etcdx)
	go func(wg *sync.WaitGroup, x *Etcdx) {
		for i := 0; i < 15; i++ {
			if i >= 6 && i <= 8 {
				x.Delete("name")
				x.Delete("age")
			}
			time.Sleep(time.Second * 1)
		}
		wg.Done()
	}(wg, etcdx)
	wg.Wait()

	time.Sleep(time.Second * 2)

	fmt.Println("-----------------------")
	name, _ := etcdx.Get("name")
	age, _ := etcdx.Get("age")
	fmt.Printf("[FINAL] name = %s, age = %s\n", name, age)
	time.Sleep(time.Second * 1)
}

func printChange(action int32, k, nv, ov string) error {
	fmt.Printf("%s [%s] action=%d %s => %s\n", time.Now().Format("2006-01-02 15:04:05"), k, action, ov, nv)
	return nil
}
