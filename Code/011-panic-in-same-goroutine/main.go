package main

import (
	"log"
	"time"
)

func main() {
	// recover should be placed into a defer function
	defer func() {
		if e := recover(); e != nil {
			log.Println("recovered from err : ", e)
			return
		}
	}()

	log.Println("now is ", time.Now().String())
	panic("a panic you not covered")
	log.Println("you probably can not see this")
}
