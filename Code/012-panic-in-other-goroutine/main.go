package main

import (
	"log"
	"time"
)

func main() {
	// recover can not handle panic in other goroutine
	defer func() {
		if e := recover(); e != nil {
			log.Println("recovered from err : ", e)
			return
		}
	}()

	log.Println("now is ", time.Now().String())

	go func() {
		panic("a panic you not covered")
		log.Println("can you see this ?")
	}()

	time.Sleep(time.Second * 5)
}
