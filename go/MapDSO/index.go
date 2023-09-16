package main

import "fmt"

func main() {
	cars := map[string]string{
		"volvo": "Volvo",
		"kia":   "Kia Forte",
	}

	if v, ok := cars["vinFast"]; ok {
		fmt.Println("Key has found in map: " + v)
	} else {
		fmt.Println("Key has not found in map!")
	}
	fmt.Println(cars)
}
