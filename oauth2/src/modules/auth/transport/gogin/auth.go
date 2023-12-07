package gogin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouteConfig(router *gin.Engine, db *gorm.DB) {
	auth := router.Group("/api/v1/auth")
	{
		auth.GET("/", Home(db))
		auth.GET("/sign-up", SignUp(db))
		auth.GET("/callback", CallBack(db))
	}
}
