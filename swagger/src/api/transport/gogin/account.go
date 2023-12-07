package gogin

import (
	"github.com/gin-gonic/gin"
)

func AccountsRouteConfig(router *gin.Engine) {
	accounts := router.Group("/api/v1/accounts")
	{
		accounts.POST("/", SignUp())
	}
}
