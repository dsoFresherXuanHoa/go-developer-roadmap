package main

import "fmt"

type Runnable interface {
	run()
}

type Vehicle struct {
	name  string
	speed int
}

func (v Vehicle) run() {
	fmt.Println("Vehicle is running....")
}

func main() {
	volvo := Vehicle{"Volvo", 200};
	volvo.run()
}