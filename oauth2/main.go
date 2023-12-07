package main

import (
	"go-auth2/src/configs"
	"go-auth2/src/modules/auth/entity"
)

func main() {
	if db, err := configs.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		const port = 3000

		entities := []interface{}{&entity.User{}}
		db.AutoMigrate(entities...)
		RouteConfig(db, port)
	}
}
