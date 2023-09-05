package main

import "fmt"

func main()  {
	var gender = 0;

	if gender == 0 {
		fmt.Println("Hi! Mr!")
	} else if gender == 1 {
		fmt.Println("Hi! Mrs!")
	} else {
		fmt.Println("Hi!")
	}

	if isError := false; !isError {
		fmt.Println(isError)
	} else {
		fmt.Println("OK!")
	}
}