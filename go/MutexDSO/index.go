package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0
	lock := new(sync.Mutex)
	for i := 1; i <= 5; i++ {
		go func() {
			for j := 1; j <= 10000; j++ {
				lock.Lock()
				count++
				fmt.Println(count)
				lock.Unlock()
			}
		}()
	}

	time.Sleep(time.Second * 3)
}
