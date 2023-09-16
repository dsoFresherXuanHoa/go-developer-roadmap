package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Index: ", i)
	}

	i := 0
	for i < 10 {
		fmt.Println("Index (while): ", i)
		i++
	}
}
