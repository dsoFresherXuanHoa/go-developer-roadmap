package main

import "fmt"

func main() {
	var cars = []string{"Volvo", "Kia Forte", "Ford Ranger"}
	var carsArchive = [...]string{"Volvo", "Kia Forte", "Ford Ranger"}
	
	carsBak := carsArchive;
	anotherCarsBak := cars;

	carsArchive[0] = "Volvo Extreme"
	cars[0] = "Volvo Extreme"

	fmt.Println(cars, carsArchive)

	fmt.Println(carsArchive, carsBak)
	fmt.Println(anotherCarsBak, cars)

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}

	for _, v := range cars {
		fmt.Println(v)
	}

	fmt.Println(len(cars), cap(cars))
}