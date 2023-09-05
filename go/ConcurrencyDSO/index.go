package main

import (
	"fmt"
	"time"
)

func pinger(msg chan string) {
	msg <- "Ping"
}

func printer(msg chan string) {
	message := <-msg
	fmt.Println(message)
	time.Sleep(time.Second * 1)
}

func main() {
	var msg chan string = make(chan string);

	go pinger(msg)
	go printer(msg)

	var input string;
	fmt.Scanln(&input)
}