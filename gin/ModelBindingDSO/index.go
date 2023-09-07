package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/login", func(ctx *gin.Context) {
		var account Account
		if err := ctx.ShouldBindJSON(&account); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"payload": account,
			})
		}
	})

	router.Run()
}
