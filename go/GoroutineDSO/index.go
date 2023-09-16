package main

import (
	"fmt"
	"time"
)

func saySomething(msg string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(i, ": ", msg)
	}
}

func main() {
	go saySomething("dsoFresherXuanHoa")
	go saySomething("dsoInternXuanHoa")

	time.Sleep(time.Second * 5)
}
