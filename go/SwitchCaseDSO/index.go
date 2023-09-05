package main

import "fmt"

func main()  {
	var gender = 0;

	switch gender {
	case 0:
		fmt.Println("Hi! Mr!")
		fallthrough
	case 1: 
		fmt.Println("Hi! Mrs!")
	case 2:
		fmt.Println("Hi!")
	default:
		fmt.Println("Finish!!!")
	} 
} 