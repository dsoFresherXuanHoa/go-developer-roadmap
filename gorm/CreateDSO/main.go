package main

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Name  string
	Speed int
}

func (v *Vehicle) BeforeCreate(tx *gorm.DB) (err error) {
	v.Name = strings.ToUpper(v.Name)
	if v.Speed > 300 {
		return errors.New("INVALID VALUE")
	}
	return
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
			result := database.Create(&vehicle)
		*/

		vehicle := Vehicle{Name: "Kia Forte", Speed: 250}
		result := database.Select("Name").Create(&vehicle)
		fmt.Println(result.Error, result.RowsAffected)

	}
}
