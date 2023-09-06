package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Id          int
	Name        string
	Speed       int
	PersonRefer int
	Person      Person `gorm:"foreignKey:PersonRefer"`
	// PersonID int

}

type Person struct {
	Id   int
	Name string
}

func main() {
	var dns = "root:PBRAGJDejNumgLLyNM4wzFYkq8JA5QXS@tcp(127.0.0.1:3306)/gorm_dev?charset=utf8mb4&parseTime=True&loc=Local"
	if database, err := gorm.Open(mysql.Open(dns), &gorm.Config{}); err != nil {
		fmt.Println("Something went wrong: \n" + err.Error())
		return
	} else {
		fmt.Println("Connection to database has been created!")

		var vehicle Vehicle
		var person Person

		database.AutoMigrate(&vehicle)
		database.AutoMigrate(&person)
	}
}
