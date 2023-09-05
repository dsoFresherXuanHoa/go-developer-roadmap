package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func goRouteTwo() {
	time.Sleep(time.Second * 2)
	fmt.Println("Goroutines - Two")
	wg.Done()
}


func goRouteOne() {
	time.Sleep(time.Second)
	fmt.Println("Goroutines - One")
	wg.Done()
}

var wg sync.WaitGroup;

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(1)
	go goRouteTwo();
	wg.Wait()

	wg.Add(1)
	go goRouteOne();
	wg.Wait()

	time.Sleep(time.Second * 5)
}