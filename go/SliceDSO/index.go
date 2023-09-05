package main

import (
	"fmt"
	"reflect"
)

func main() {
	var cars = [...]string{"Volvo", "Kia Forte", "Ford Ranger"}
	var carsSlice = cars[0:]

	fmt.Println(cars, reflect.TypeOf(cars))
	fmt.Println(carsSlice, reflect.TypeOf(carsSlice))

	fmt.Println(len(carsSlice), cap(carsSlice))

	var vehicles =  make([]int, 3, 5);
	var vehiclesBak = make([]int, 3, 5);
	var subVehicles = vehicles[0:len(vehicles) - 1]

	copy(vehicles, vehiclesBak)

	vehicles = append(vehicles, 1)
	vehicles = append(vehicles, 1)
	vehicles = append(vehicles, 1)
	fmt.Println(len(vehicles), cap(vehicles))

	fmt.Println(vehicles)
	fmt.Println(vehiclesBak)
	fmt.Println(subVehicles)
}