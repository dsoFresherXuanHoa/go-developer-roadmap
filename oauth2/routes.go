package main

import (
	"os"
	"path/filepath"
	"strconv"

	authGin "go-auth2/src/modules/auth/transport/gogin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteConfig(db *gorm.DB, port int) {
	currentPath, _ := os.Getwd()
	viewsPath := filepath.Join(currentPath, "./src/static/views/*")

	router := gin.Default()
	router.LoadHTMLGlob(viewsPath)
	authGin.AuthRouteConfig(router, db)
	router.Run(":" + strconv.Itoa(port))
}
