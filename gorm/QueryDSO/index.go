package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Name  string
	Speed int
}

type ResponseVehicle struct {
	gorm.Model
	Name string
}

func main() {
	var dns = "root:PBRAGJDejNumgLLyNM4wzFYkq8JA5QXS@tcp(127.0.0.1:3306)/gorm_dev?charset=utf8mb4&parseTime=True&loc=Local"
	if database, err := gorm.Open(mysql.Open(dns), &gorm.Config{}); err != nil {
		fmt.Println("Something went wrong: " + err.Error())
		return
	} else {
		fmt.Println("Connection to database has been created!")
		// fmt.Println(database)

		var vehicle Vehicle
		var responseVehicle ResponseVehicle

		// firstVehicle := database.First(&vehicle)
		// lastVehicle := database.Last(&vehicle)
		// secondVehicle := database.First(&vehicle, 1)
		// someVehicle := database.Find(&vehicle, []int{1, 2, 3})

		// fmt.Println(firstVehicle)
		// fmt.Println(lastVehicle)
		// fmt.Println(secondVehicle)
		// fmt.Println(someVehicle.Rows())

		// fmt.Println(database.Where("id = ?", 2).Find(&vehicle))
		// fmt.Println(database.Where("id IN ?", []int{1, 2}).Find(&vehicle))

		// fmt.Println(database.Order("speed desc").Find(&vehicle))
		// fmt.Println(database.Limit(2).Offset(3).Find(&vehicle))
		// fmt.Println(database.Distinct("speed").Find(&vehicle))
		// fmt.Println(database.Model(&vehicle).Where("id = ?", 1).Find(&responseVehicle))
		// fmt.Println(responseVehicle)

		// database.First(&vehicle)
		// vehicle.Name = "Volvo"
		// database.Save(&vehicle)

		database.Model(&vehicle).Where("id > ?", 10).Find(&responseVehicle)
		fmt.Println(vehicle)
		fmt.Println(responseVehicle)
	}
}
