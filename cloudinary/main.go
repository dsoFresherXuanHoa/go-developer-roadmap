package main

import (
	"go-cloudinary/src/config"
	"go-cloudinary/src/src/module/entity"
)

func main() {
	if db, err := config.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else if cld, err := config.GetCloudinaryInstance(); err != nil {
		panic("Can't connect to cloudinary server: " + err.Error())
	} else {
		const port = 3000
		entities := []interface{}{
			&entity.Thumbs{},
		}
		db.AutoMigrate(entities...)
		RouteConfig(db, cld, port)
	}
}
