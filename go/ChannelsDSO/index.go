package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)

	go func() {
		msg <- "dsoFresherXuanHoa"
		fmt.Println("Next!")
		}()
		
	fmt.Println(<-msg)
	time.Sleep(time.Second * 3)
}