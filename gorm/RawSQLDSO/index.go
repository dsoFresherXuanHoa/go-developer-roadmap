package main

import (
	"fmt"
	"reflect"

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

		var vehicle Vehicle

		database.Raw("SELECT * FROM vehicles").Scan(&vehicle)
		fmt.Println(reflect.TypeOf(vehicle))
	}
}
