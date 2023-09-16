package main

import (
	"fmt"
	"time"
)

func sum(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go sum(numbers[:len(numbers)/2], ch)
	go sum(numbers[len(numbers)/2:], ch)

	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

	// Buffer Chanel
	bufferChanel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		bufferChanel <- i
	}
	close(bufferChanel)

	for i := range bufferChanel {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}

	for i := 0; i < 15; i++ {
		if chanelValue, ok := <-bufferChanel; ok {
			fmt.Println(chanelValue)
		} else {
			fmt.Println("BREAKING...")
			break
		}
	}

	// Select Chanel
	chanelOne := make(chan int)
	chanelTwo := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		chanelTwo <- 2
	}()
	go func() {
		time.Sleep(time.Second)
		chanelOne <- 1
	}()

	select {
	case one := <-chanelOne:
		fmt.Println(one)
	case two := <-chanelTwo:
		fmt.Println(two)
	}

	time.Sleep(time.Second * 3)
}
