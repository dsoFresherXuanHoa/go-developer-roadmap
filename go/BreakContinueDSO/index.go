package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if (i == 10) {
			fmt.Println("Skipped!");
			continue;
		}
		if (i == 20) {
			fmt.Println("Exited!")
			break
		}
		fmt.Println(i)
	}
}