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

func main() {
	var dns = "root:PBRAGJDejNumgLLyNM4wzFYkq8JA5QXS@tcp(127.0.0.1:3306)/gorm_dev?charset=utf8mb4&parseTime=True&loc=Local"
	if database, err := gorm.Open(mysql.Open(dns), &gorm.Config{}); err != nil {
		fmt.Println("Something went wrong: " + err.Error())
		return
	} else {
		fmt.Println("Connection to database has been created!")
		fmt.Println(database)

		/*
			database.AutoMigrate(&Vehicle{})
			fmt.Println("Schema has been migrated!")

			database.Create(&Vehicle{Name: "Volvo", Speed: 200})
			fmt.Println("Record has been saved!")

			firstVehicle := database.First(&vehicle, 1)
			fmt.Println(firstVehicle)

			secondVehicle := database.First(&vehicle, "id = ?", 1)
			fmt.Println(secondVehicle)
			database.Model(&vehicle).Update("Name", "Volvo Extreme")

			var vehicle Vehicle
			database.Delete(&vehicle, 2)
		*/
	}
}
