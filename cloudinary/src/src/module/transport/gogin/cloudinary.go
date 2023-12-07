package gogin

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CloudinaryRouteConfig(router *gin.Engine, db *gorm.DB, cld *cloudinary.Cloudinary) {
	cloudinary := router.Group("/api/v1/cloudinary")
	{
		cloudinary.POST("/upload", Upload())
	}
}
