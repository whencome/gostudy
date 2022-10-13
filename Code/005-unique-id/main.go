package main

import (
	"fmt"
	"sync"

	ulid "github.com/oklog/ulid/v2"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(v int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				id := ulid.Make()
				fmt.Printf("%d => %s\n", v, id.String())
			}
		}(i)
	}
	wg.Wait()
}
