package main

import (
	"go-cloudinary/src/src/module/transport/gogin"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteConfig(db *gorm.DB, cld *cloudinary.Cloudinary, port int) {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	gogin.CloudinaryRouteConfig(router, db, cld)
	router.Run(":" + strconv.Itoa(port))
}
