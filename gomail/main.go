package main

import (
	"fmt"
	"go-gomail/src/config"
	"go-gomail/src/module/entity"
)

func main() {
	if db, err := config.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		const port = 3000
		entities := []interface{}{
			&entity.UserCreatable{},
		}

		db.AutoMigrate(entities...)
		fmt.Println("All entity has been synced to database!")
		RouteConfig(db, port)
	}
}
