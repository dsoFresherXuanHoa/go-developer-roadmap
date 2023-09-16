package main

import "fmt"

func main() {
	cars := [3]string{"Volvo", "Kia Forte", "Ford Ranger"}
	carsSlice := cars[:]
	carsMake := make([]string, 0, 5)

	fmt.Println(carsSlice)
	cars[0] = "Toyota"
	carsSlice = append(carsSlice, "VinFast")

	fmt.Println(carsSlice)
	fmt.Println(len(carsSlice), cap(carsSlice))
	fmt.Println(len(carsMake), cap(carsMake))

	for i, v := range carsSlice {
		fmt.Println("Index: ", i, ". Value: ", v)
	}
}
