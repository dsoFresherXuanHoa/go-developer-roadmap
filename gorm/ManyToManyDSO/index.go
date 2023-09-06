package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

func main() {
	var dns = "root:PBRAGJDejNumgLLyNM4wzFYkq8JA5QXS@tcp(127.0.0.1:3306)/gorm_dev?charset=utf8mb4&parseTime=True&loc=Local"
	if database, err := gorm.Open(mysql.Open(dns), &gorm.Config{}); err != nil {
		fmt.Println("Something went wrong: \n" + err.Error())
		return
	} else {
		fmt.Println("Connection to database has been created!")

		var users User
		var languages Language

		database.AutoMigrate(&users)
		database.AutoMigrate(&languages)
	}
}
