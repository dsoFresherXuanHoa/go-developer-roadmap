package main

import (
	docs "go-swagger/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go-swagger/src/api/entity"
	"go-swagger/src/api/transport/gogin"
	"go-swagger/src/config"
	"os"

	"github.com/gin-gonic/gin"
)

// @title Account Service API
func main() {
	if db, err := config.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		port := os.Getenv("PORT")
		models := []interface{}{
			&entity.Account{},
		}
		db.AutoMigrate(models...)

		router := gin.Default()
		docs.SwaggerInfo.BasePath = "/api/v1"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		router.MaxMultipartMemory = 8 << 20
		gogin.AccountsRouteConfig(router)
		router.Run(":" + port)
	}
}
