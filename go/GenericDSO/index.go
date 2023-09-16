package main

import "fmt"

func Index[T comparable](src []T, val T) bool {
	for _, v := range src {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	cars := []string{"Volvo", "Kia Forte"}
	numbers := []int{1, 2, 3}

	fmt.Println(Index[string](cars, "Volvo"))
	fmt.Println(Index[int](numbers, 1))
}
