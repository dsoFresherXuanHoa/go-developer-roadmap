package main

import "fmt"

func main() {
	var cars = map[string]string {
		"volvo": "Volvo",
		"kia": "Kia Forte",
	}

	cars["ford"] = "Ford Ranger"

	delete(cars, "ford")

	fmt.Println(len(cars))
	fmt.Println(cars)

	if volvo, ok := cars["volvo-extreme"]; ok {
		fmt.Println(volvo)
	} else {
		fmt.Println("Key not found in map!")
	}
}