package main

import (
	"fmt"
	"reflect"
)

type Vehicle struct {
	name  string
	speed int
}

func (v Vehicle) getName() string {
	return v.name
}

func (v *Vehicle) setName(name string) {
	v.name = name;
}

func main() {
	volvo := Vehicle{"Volvo", 200}

	fmt.Println(volvo)
	fmt.Println(reflect.TypeOf(volvo))

	pVolvo := &volvo;

	fmt.Println(pVolvo.name)

	// Receiver
	fmt.Println(volvo)
	volvo.setName("Volvo Extreme")
	fmt.Println(volvo)
}