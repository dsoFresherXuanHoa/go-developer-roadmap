package main

import (
	"fmt"
	"strings"
)

func args(args ...string) {
	fmt.Println(args)
}

func toUpperCase(cars []string)  {
	for i, v := range cars {
		cars[i] = strings.ToUpper(v);
	}
}

func main()  {
	var cars = []string{"Volvo", "Kia Forte", "Ford Ranger"}

	args("dsoFresherXuanHoa", "dsoFresherXuanHoa");
	args("dsoFresherXuanHoa"); 

	toUpperCase(cars)
	fmt.Println(cars)
}