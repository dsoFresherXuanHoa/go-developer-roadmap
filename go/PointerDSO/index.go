package main

import "fmt"

func main() {
	age := 22
	pAge := &age

	fmt.Println(age, pAge, *pAge)
	*pAge = 23

	fmt.Println(age, pAge, *pAge)
}
